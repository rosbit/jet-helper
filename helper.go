package tmpl

import (
	"github.com/CloudyKit/jet/v6"
	"io"
)

type (
	Template struct {
		*jet.Template
		*jet.Set
		vars jet.VarMap
	}

	Tmpls struct {
		*jet.Set
	}

	FuncMap map[string]interface{}

	Option = jet.Option
	SafeWriter = jet.SafeWriter
	Cache = jet.Cache
)

var (
	InDevelopmentMode = jet.InDevelopmentMode
	WithDelims = jet.WithDelims
	WithSafeWriter = jet.WithSafeWriter
	WithCache = jet.WithCache
)

func NewTempls(homeDir string, global ...FuncMap) *Tmpls {
	set := jet.NewSet(jet.NewOSFileSystemLoader(homeDir))
	if len(global) > 0 && len(global[0]) > 0 {
		for n, v := range global[0] {
			set.AddGlobal(n, v)
		}
	}
	return &Tmpls {
		Set: set,
	}
}

func (ts *Tmpls) Lookup(name string) (*Template, error) {
	t, err := ts.Set.GetTemplate(name)
	if err != nil {
		return nil, err
	}
	return &Template{
		Template: t,
		Set: ts.Set,
	}, nil
}

func (ts *Tmpls) Delims(left, right string) *Tmpls {
	opt := jet.WithDelims(left, right)
	opt(ts.Set)
	return ts
}

func (ts *Tmpls) Options(opts ...jet.Option) *Tmpls {
	for _, opt := range opts {
		opt(ts.Set)
	}
	return ts
}

func (ts *Tmpls) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	t, err := ts.Lookup(name)
	if err != nil {
		return err
	}
	return t.Execute(wr, data)
}

func Parse(text string) (*Template, error) {
	const tmpl_name = "jet"

	loader := jet.NewInMemLoader()
	loader.Set(tmpl_name, text)
	set := jet.NewSet(loader)
	t, err := set.GetTemplate(tmpl_name)
	if err != nil {
		return nil, err
	}
	return &Template{
		Template: t,
		Set: set,
	}, nil
}

func (t *Template) Funcs(funcMap FuncMap) *Template {
	t.vars = make(jet.VarMap)
	for n, v := range funcMap {
		t.vars.Set(n, v)
	}
	return t
}

func (t *Template) Execute(wr io.Writer, data interface{}) error {
	return t.Template.Execute(wr, t.vars, data)
}
