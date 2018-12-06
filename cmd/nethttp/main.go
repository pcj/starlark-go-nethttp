package main // import "github.com/pcj/starlark-go-nethttp/cmd/nethttp"

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"sort"
	"strings"

	"github.com/pcj/starlark-go-nethttp"
	
	"go.starlark.net/repl"
	"go.starlark.net/resolve"
	"go.starlark.net/starlark"
)

// flags
var (
	cpuprofile = flag.String("cpuprofile", "", "gather CPU profile in this file")
	showenv    = flag.Bool("showenv", false, "on success, print final global environment")
)

// non-standard dialect flags
func init() {
	flag.BoolVar(&resolve.AllowFloat, "fp", resolve.AllowFloat, "allow floating-point numbers")
	flag.BoolVar(&resolve.AllowSet, "set", resolve.AllowSet, "allow set data type")
	flag.BoolVar(&resolve.AllowLambda, "lambda", resolve.AllowLambda, "allow lambda expressions")
	flag.BoolVar(&resolve.AllowNestedDef, "nesteddef", resolve.AllowNestedDef, "allow nested def statements")
	flag.BoolVar(&resolve.AllowBitwise, "bitwise", resolve.AllowBitwise, "allow bitwise operations (&, |, ^, ~, <<, and >>)")
}

func main() {
	log.SetPrefix("nethttp: ")
	log.SetFlags(0)
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
	}

	thread := &starlark.Thread{Load: repl.MakeLoad()}
	globals := starlark.StringDict{
        "http": nethttp.NewModule(),
    }

	switch len(flag.Args()) {
	case 0:
		fmt.Println("Welcome to nethttp")
		repl.REPL(thread, globals)
	case 1:
		// Execute specified file.
		filename := flag.Args()[0]
		var err error
		globals, err = starlark.ExecFile(thread, filename, nil, nil)
		if err != nil {
			repl.PrintError(err)
			os.Exit(1)
		}
	default:
		log.Fatal("want at most one Starlark file name")
	}

	// Print the global environment.
	if *showenv {
		var names []string
		for name := range globals {
			if !strings.HasPrefix(name, "_") {
				names = append(names, name)
			}
		}
		sort.Strings(names)
		for _, name := range names {
			fmt.Fprintf(os.Stderr, "%s = %s\n", name, globals[name])
		}
	}
}
