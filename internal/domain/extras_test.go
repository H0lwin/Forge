package domain

import "testing"

func TestValidateFrameworkExtras(t *testing.T) {
	if err := ValidateFrameworkExtras(FrameworkDjango, []string{"drf"}); err != nil {
		t.Fatalf("expected drf to be valid for django: %v", err)
	}
	if err := ValidateFrameworkExtras(FrameworkNext, []string{"drf"}); err == nil {
		t.Fatalf("expected drf invalid for next")
	}
	if err := ValidateFrameworkExtras(FrameworkExpress, []string{"tailwind"}); err == nil {
		t.Fatalf("expected tailwind invalid for express")
	}
}
