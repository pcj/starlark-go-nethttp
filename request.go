package nethttp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"go.starlark.net/starlark"
)

func fnHttpRequest(t *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	request, err := newRequestFromStarlarkArgs("GET", args, kwargs)
	if err != nil {
		return nil, err
	}
	return NewHttpRequest(request)
}

func NewHttpRequest(request *http.Request) (*HttpRequest, error) {
	attrs, err := GetHttpRequestAttrs(request)
	if err != nil {
		return nil, err
	}
	return &HttpRequest{
		request: request,
		attrs:   *attrs,
	}, nil
}

type HttpRequest struct {
	request *http.Request
	attrs   starlark.StringDict
}

var _ starlark.HasAttrs = (*HttpRequest)(nil)

func (r *HttpRequest) String() string {
	return fmt.Sprintf("Request(%s)", r.attrs.String())
}
func (r *HttpRequest) Type() string         { return "module" }
func (r *HttpRequest) Freeze()              { r.attrs.Freeze() }
func (r *HttpRequest) Truth() starlark.Bool { return starlark.True }
func (r *HttpRequest) Hash() (uint32, error) {
	return 0, fmt.Errorf("unhashable type: %s", r.Type())
}

func (r *HttpRequest) Attr(name string) (starlark.Value, error) {
	if val, ok := r.attrs[name]; ok {
		return val, nil
	}
	return nil, nil
}

func (r *HttpRequest) AttrNames() []string {
	var names []string
	for name := range r.attrs {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func GetHttpRequestAttrs(request *http.Request) (*starlark.StringDict, error) {
	bytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}
	return &starlark.StringDict{
		"method":         starlark.String(request.Method),
		"url":            starlark.String(request.URL.String()),
		"host":           starlark.String(request.Host),
		"body":           starlark.String(string(bytes)),
		"content_length": starlark.MakeInt(int(request.ContentLength)),
		"headers":        getDictFromHeaders(request.Header),
	}, nil
}

func newRequestFromStarlarkArgs(method string, args starlark.Tuple, kwargs []starlark.Tuple) (*http.Request, error) {

	request := &http.Request{
		Method: method,
	}

	var rawurl, body string
	headers := &starlark.Dict{}
	params := &starlark.Dict{}

	if err := starlark.UnpackArgs("http.Request", args, kwargs,
		"url", &rawurl,
		"method?", &request.Method,
		"host?", &request.Host,
		"body?", &body,
		"headers?", &headers,
		"params?", &params,
	); err != nil {
		return nil, err
	}

	parsedURL, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	param, err := getMapFromDict(params)
	if err != nil {
		return nil, err
	}

	q := parsedURL.Query()
	for k, v := range param {
		q[k] = append(q[k], v...)
	}
	parsedURL.RawQuery = q.Encode()

	header, err := getMapFromDict(headers)
	if err != nil {
		return nil, err
	}

	request.URL = parsedURL
	request.Body = ioutil.NopCloser(strings.NewReader(body))
	request.ContentLength = int64(len(body))
	request.Header = header

	return request, nil
}

func getMapFromDict(dict *starlark.Dict) (map[string][]string, error) {
	m := make(map[string][]string, dict.Len())
	for _, item := range dict.Items() {
		k, ok := item[0].(starlark.String)
		if !ok {
			return nil, fmt.Errorf("Expected string key type for: %+v", item)
		}
		key := k.GoString()

		switch t := item[1].(type) {
		case starlark.String:
			m[key] = append(m[key], t.GoString())
		case *starlark.List:
			m[key] = getStringSliceFromStarlarkList(t)
		}
	}
	return m, nil
}

func getStringSliceFromStarlarkList(list *starlark.List) []string {
	size := list.Len()
	slice := make([]string, size)
	for i := 0; i < size; i++ {
		value := list.Index(i)
		var str string
		switch t := value.(type) {
		case starlark.String:
			str = t.GoString()
		default:
			str = t.String()
		}
		slice[i] = str
	}
	return slice
}

func getDictFromHeaders(headers http.Header) *starlark.Dict {
	dict := &starlark.Dict{}
	for key, values := range headers {
		dict.SetKey(starlark.String(key), getSkylarkList(values))
	}
	return dict
}

func getSkylarkList(values []string) *starlark.List {
	list := &starlark.List{}
	for _, v := range values {
		list.Append(starlark.String(v))
	}
	return list
}

// AsHttpRequest attempts a type assertion and returns the typed request or False
// if the type assertion failed.
func AsHttpRequest(value starlark.Value) (*HttpRequest, bool) {
	if req, ok := value.(*HttpRequest); ok {
		return req, true
	}
	return nil, false
}
