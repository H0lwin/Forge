package config

import (
	"path/filepath"
	"testing"
)

func TestSaveLoadRoundtrip(t *testing.T) {
	d := t.TempDir()
	path := filepath.Join(d, "config.yaml")
	cfg := Default()
	cfg.User.Name = "Tester"
	if err := Save(path, cfg); err != nil {
		t.Fatalf("save: %v", err)
	}
	got, err := Load(path)
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	if got.User.Name != "Tester" {
		t.Fatalf("unexpected user name: %q", got.User.Name)
	}
}
