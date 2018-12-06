package nethttp

import (
	"io/ioutil"
	"net/http"

	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
)

// FromHttpResponse constructs a new starlark StringDict from an http.Response
// and error value.
func newHttpResponse(resp *http.Response, err error) (starlark.Value, error) {

	// Developer should always be able to access all properties on the response
	// struct, possibly set to the "zero value" for that type.

	var body, status, errMessage string
	var code int

	if err != nil {
		errMessage = err.Error()
	}

	if resp != nil {
		defer resp.Body.Close()
		bytes, err := ioutil.ReadAll(resp.Body)
		// TODO: should this error be embedded in the response as well?
		if err != nil {
			return nil, err
		}
		body = string(bytes)
		status = resp.Status
		code = resp.StatusCode
	}

	return starlarkstruct.FromStringDict(
		starlark.String("http.response"),
		starlark.StringDict{
			"body":   starlark.String(body),
			"status": starlark.String(status),
			"code":   starlark.MakeInt(code),
			"error":  starlark.String(errMessage),
		},
	), nil
}
