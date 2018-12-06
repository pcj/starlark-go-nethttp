package nethttp

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"

	"go.starlark.net/starlark"
)

func TestModule(t *testing.T) {

	tests := []testCase{
		{
			name:   "module should identify itself",
			input:  `print(http)`,
			output: `<module http>`,
		},
	}

	for i, tc := range tests {
		runTestCase(t, i, &tc)
	}
}

type testCase struct {
	name   string
	input  string
	output string
}

func runTestCase(t *testing.T, i int, tc *testCase) {
	t.Run(tc.name, func(t *testing.T) {
		var out bytes.Buffer
		thread := &starlark.Thread{
			Print: func(_ *starlark.Thread, msg string) { fmt.Fprintln(&out, msg) },
		}
		globals := starlark.StringDict{
			"http": NewModule(),
		}

		_, err := starlark.ExecFile(thread, tc.name, tc.input, globals)

		if err != nil {
			if evalErr, ok := err.(*starlark.EvalError); ok {
				log.Fatal(evalErr.Backtrace())
			}
			log.Fatal(err)
		}

		got := strings.TrimSpace(out.String())
		want := tc.output

		if got != want {
			t.Errorf("#%d: got %q, want %q", i, got, want)
		}
	})
}
