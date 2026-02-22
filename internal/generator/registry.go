package generator

import (
	"fmt"

	"forge/internal/domain"
)

type Registry struct {
	items map[domain.Framework]Generator
}

func NewRegistry(gens ...Generator) *Registry {
	m := map[domain.Framework]Generator{}
	for _, g := range gens {
		m[domain.Framework(g.Name())] = g
	}
	return &Registry{items: m}
}

func (r *Registry) Get(framework domain.Framework) (Generator, error) {
	g, ok := r.items[framework]
	if !ok {
		return nil, fmt.Errorf("framework not supported: %s", framework)
	}
	return g, nil
}

func (r *Registry) List() []string {
	out := make([]string, 0, len(r.items))
	for _, g := range r.items {
		out = append(out, g.Name())
	}
	return out
}
