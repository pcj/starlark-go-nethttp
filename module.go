package nethttp

import (
	"fmt"
	"sort"

	"go.starlark.net/starlark"
)

// NewModule constructs the module instance that can be
// added to the global scope.
func NewModule() *Module {
	mod := &Module{
		attrs: starlark.StringDict{
			"do":      starlark.NewBuiltin("http.do", fnHttpDo),
			"get":     starlark.NewBuiltin("http.get", fnHttpGet),
			"patch":   starlark.NewBuiltin("http.patch", fnHttpPatch),
			"post":    starlark.NewBuiltin("http.post", fnHttpPost),
			"put":     starlark.NewBuiltin("http.put", fnHttpPut),
			"delete":  starlark.NewBuiltin("http.delete", fnHttpDelete),
			"Request": starlark.NewBuiltin("http.Request", fnHttpRequest),
		},
	}
	return mod
}

type Module struct {
	attrs starlark.StringDict
}

var _ starlark.HasAttrs = (*Module)(nil)

func (mod *Module) String() string       { return "<module http>" }
func (mod *Module) Type() string         { return "module" }
func (mod *Module) Freeze()              { mod.attrs.Freeze() }
func (mod *Module) Truth() starlark.Bool { return starlark.True }
func (mod *Module) Hash() (uint32, error) {
	return 0, fmt.Errorf("unhashable type: %s", mod.Type())
}

func (mod *Module) Attr(name string) (starlark.Value, error) {
	if val, ok := mod.attrs[name]; ok {
		return val, nil
	}
	return nil, nil
}

func (mod *Module) AttrNames() []string {
	var names []string
	for name := range mod.attrs {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
