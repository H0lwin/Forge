package system

import "testing"

func TestParseVersion(t *testing.T) {
	if got := ParseVersion("git version 2.43.0"); got != "2.43.0" {
		t.Fatalf("got %q", got)
	}
	if got := ParseVersion("no version"); got != "-" {
		t.Fatalf("got %q", got)
	}
}
