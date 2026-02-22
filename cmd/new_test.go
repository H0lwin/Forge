package cmd

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestNewNoInteractiveMissingFlags(t *testing.T) {
	var out bytes.Buffer
	root := NewRootCommand(context.Background(), &out, &out)
	root.SetArgs([]string{"new", "--no-interactive", "--framework", "django"})
	err := root.Execute()
	if err == nil {
		t.Fatalf("expected error")
	}
	if !strings.Contains(err.Error(), "missing required flag") {
		t.Fatalf("unexpected error: %v", err)
	}
}
