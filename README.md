# starlark-go-nethttp

A wrapper around a minimal subset of `net/http` package for use within [starlark](https://github.com/google/starlark-go).

## Installation

Include the module and add it to your global scope:

```go
package main 

import (
    "flag"
    "fmt"
    "os"

    "github.com/pcj/starlark-go-nethttp"

    "go.starlark.net/repl"
    "go.starlark.net/starlark"
)

func main() {
    thread := &starlark.Thread{
        Print: func(_ *starlark.Thread, msg string) { fmt.Println(msg) },
    }
    globals := &starlark.StringDict{
        "http": nethttp.NewModule(),
    }
    filename := flag.Args()[0]
    var err error
    globals, err = starlark.ExecFile(thread, filename, nil, globals)
    if err != nil {
        repl.PrintError(err)
        os.Exit(1)
    }
}
```

## Usage 

Then, within your starlark script:

```python
response = http.get("https//example.com")

print("{status} ({code}): error={error}\n\n{body}".format(
    status = response.status,
    code = response.code,
    error = response.error,
    body = response.body,
))

```
