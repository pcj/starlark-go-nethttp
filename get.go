package nethttp

import (
	"net/http"

	"go.starlark.net/starlark"
)

// Implementation of the `http.get()` built-in function.
// Reset protobuf state to the default values.
func fnHttpGet(t *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {

	request, err := newRequestFromStarlarkArgs("GET", args, kwargs)
	if err != nil {
		return nil, err
	}

	return newHttpResponse(http.DefaultClient.Do(request))
}
