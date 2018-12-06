# starlark-go-nethttp

A wrapper around a minimal subset of `net/http` package for use within
[starlark-go](https://github.com/google/starlark-go).

## Documentation

* API documentation: [godoc.org/github.com/pcj/starlark-go-nethttp](https://godoc.org/github.com/pcj/starlark-go-nethttp)

### Getting started

Build the code:

```shell
# check out the code and dependencies,
# and install interpreter in $GOPATH/bin
$ go get -u github.com/pcj/starlark-go-nethttp
```

Run the interpreter or interact with the read-eval-print loop (REPL):

```
$ nethttp
>>> resp = http.get("https://google.com")
>>> resp.code
200
>>>
```

When you have finished, type `Ctrl-D` to close the REPL's input stream. 

### Embedding

To embed the module within your own configuration language, add it to your globals:

```python
globals := starlark.StringDict{
    "http": nethttp.NewModule(),
}
```

### Contributing

Contributions welcome.
