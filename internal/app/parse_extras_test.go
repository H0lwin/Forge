package app

import (
	"reflect"
	"testing"
)

func TestParseExtrasDeduplicateAndSort(t *testing.T) {
	got := ParseExtras("docker, ci,DRF,docker,ci")
	want := []string{"ci", "docker", "drf"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
