package templates

import (
	"embed"
	"fmt"
	"path"
	"strings"
	"text/template"
)

//go:embed *.tmpl
var FS embed.FS

type Engine struct {
	root *template.Template
}

func NewEngine() (*Engine, error) {
	files, err := FS.ReadDir(".")
	if err != nil {
		return nil, err
	}
	t := template.New("root").Option("missingkey=error")
	for _, f := range files {
		if f.IsDir() || !strings.HasSuffix(f.Name(), ".tmpl") {
			continue
		}
		b, readErr := FS.ReadFile(f.Name())
		if readErr != nil {
			return nil, readErr
		}
		_, parseErr := t.New(f.Name()).Parse(string(b))
		if parseErr != nil {
			return nil, parseErr
		}
	}
	return &Engine{root: t}, nil
}

func (e *Engine) Render(name string, data any) ([]byte, error) {
	var b strings.Builder
	t := e.root.Lookup(name)
	if t == nil {
		return nil, fmt.Errorf("template not found: %s", name)
	}
	if err := t.Execute(&b, data); err != nil {
		return nil, err
	}
	return []byte(b.String()), nil
}

func (e *Engine) List() []string {
	out := []string{}
	for _, t := range e.root.Templates() {
		if t.Name() == "root" {
			continue
		}
		out = append(out, path.Base(t.Name()))
	}
	return out
}
