package domain

import "testing"

func TestMissingRequiredFlags_NoInteractiveDjango(t *testing.T) {
	req := GenerateRequest{Framework: FrameworkDjango}
	missing := MissingRequiredFlags(req)
	want := map[string]bool{"--name": true, "--path": true, "--python-version": true, "--env-manager": true}
	for _, f := range missing {
		delete(want, f)
	}
	if len(want) != 0 {
		t.Fatalf("missing flags mismatch: %+v", want)
	}
}

func TestValidateName(t *testing.T) {
	if err := ValidateName("my-app"); err != nil {
		t.Fatalf("expected valid name: %v", err)
	}
	if err := ValidateName("Bad_Name"); err == nil {
		t.Fatalf("expected invalid name error")
	}
}
