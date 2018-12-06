package nethttp

import (
	"net/http"

	"go.starlark.net/starlark"
)

// fnHttpPost implements the `http.post()` built-in function.
func fnHttpPost(t *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {

	request, err := newRequestFromStarlarkArgs("POST", args, kwargs)
	if err != nil {
		return nil, err
	}

	return newHttpResponse(http.DefaultClient.Do(request))
}
