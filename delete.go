package nethttp

import (
	"net/http"

	"go.starlark.net/starlark"
)

// fnHttpDelete implements the `http.Delete()` built-in function.
func fnHttpDelete(t *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {

	request, err := newRequestFromStarlarkArgs("DELETE", args, kwargs)
	if err != nil {
		return nil, err
	}

	return newHttpResponse(http.DefaultClient.Do(request))
}
