package generator

import (
	"context"
	"testing"

	"forge/internal/domain"
	"forge/internal/runner"
)

type fakeGenerator struct{ name string }

func (f fakeGenerator) Name() string { return f.name }
func (f fakeGenerator) Category() string { return "backend" }
func (f fakeGenerator) Steps(context.Context, domain.GenerateRequest) []runner.Step { return nil }
func (f fakeGenerator) PreCheck(context.Context, domain.GenerateRequest) error { return nil }
func (f fakeGenerator) PostMessage(domain.GenerateRequest, domain.GenerateResult) string { return "ok" }

func TestRegistryGet(t *testing.T) {
	r := NewRegistry(fakeGenerator{name: "django"})
	if _, err := r.Get(domain.Framework("django")); err != nil {
		t.Fatalf("expected generator: %v", err)
	}
	if _, err := r.Get(domain.Framework("unknown")); err == nil {
		t.Fatalf("expected error")
	}
}
