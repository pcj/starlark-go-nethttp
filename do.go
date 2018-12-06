package nethttp

import (
	"fmt"
	"net/http"

	"go.starlark.net/starlark"
)

// fnHttpDo implements the `http.do()` built-in function.
func fnHttpDo(t *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {

	var request starlark.Value

	if err := starlark.UnpackArgs("http.Request", args, kwargs,
		"request", &request,
	); err != nil {
		return nil, err
	}

	req, ok := AsHttpRequest(request)
	if !ok {
		return nil, fmt.Errorf("Expected http.Request argument, got %q", req.String())
	}

	return newHttpResponse(http.DefaultClient.Do(req.request))
}
