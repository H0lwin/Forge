package app

import "testing"

func TestAddonSupported(t *testing.T) {
	if !AddonSupported("django", "celery") {
		t.Fatalf("expected celery supported for django")
	}
	if AddonSupported("next", "celery") {
		t.Fatalf("expected celery not supported for next")
	}
	if !AddonSupported("next", "docker") {
		t.Fatalf("expected docker supported for next")
	}
}
