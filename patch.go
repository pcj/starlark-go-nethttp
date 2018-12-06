package nethttp

import (
	"net/http"

	"go.starlark.net/starlark"
)

// fnHttpPatch implements the `http.patch()` built-in function.
func fnHttpPatch(t *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {

	request, err := newRequestFromStarlarkArgs("PATCH", args, kwargs)
	if err != nil {
		return nil, err
	}

	return newHttpResponse(http.DefaultClient.Do(request))
}
