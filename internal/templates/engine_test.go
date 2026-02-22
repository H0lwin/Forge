package templates

import "testing"

func TestRenderStrictKey(t *testing.T) {
	eng, err := NewEngine()
	if err != nil {
		t.Fatalf("engine: %v", err)
	}
	if _, err := eng.Render("README.md.tmpl", map[string]string{"Name": "x"}); err == nil {
		t.Fatalf("expected missing key error")
	}
}
