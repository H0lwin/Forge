package cmd

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestFrameworksCommand(t *testing.T) {
	var out bytes.Buffer
	root := NewRootCommand(context.Background(), &out, &out)
	root.SetArgs([]string{"frameworks"})
	if err := root.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}
	if !strings.Contains(out.String(), "django") {
		t.Fatalf("expected framework list output, got: %s", out.String())
	}
}
