package nethttp

import (
	"net/http"

	"go.starlark.net/starlark"
)

// fnHttpPut implements the `http.put()` built-in function.
func fnHttpPut(t *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {

	request, err := newRequestFromStarlarkArgs("PUT", args, kwargs)
	if err != nil {
		return nil, err
	}

	return newHttpResponse(http.DefaultClient.Do(request))
}
