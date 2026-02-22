package cmd

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestConfigMutuallyExclusiveActions(t *testing.T) {
	var out bytes.Buffer
	root := NewRootCommand(context.Background(), &out, &out)
	root.SetArgs([]string{"config", "--edit", "--reset"})
	err := root.Execute()
	if err == nil {
		t.Fatalf("expected error")
	}
	if !strings.Contains(err.Error(), "only one action") {
		t.Fatalf("unexpected error: %v", err)
	}
}
