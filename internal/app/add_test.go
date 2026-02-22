package app

import "testing"

func TestDetectFrameworkFromPath(t *testing.T) {
	d := t.TempDir()
	if got := DetectFrameworkFromPath(d); got != "" {
		t.Fatalf("expected empty framework, got %s", got)
	}
}
