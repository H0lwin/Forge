package domain

import "testing"

func TestValidateFrameworkExtrasUnsupported(t *testing.T) {
	if err := ValidateFrameworkExtras(FrameworkDjango, []string{"unknown"}); err == nil {
		t.Fatalf("expected error for unsupported extra")
	}
}
