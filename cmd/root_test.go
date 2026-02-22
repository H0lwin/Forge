package cmd

import "testing"

func TestParseFrameworkAliases(t *testing.T) {
	f, err := parseFramework("nextjs")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f != "next" {
		t.Fatalf("expected next, got %s", f)
	}
}
