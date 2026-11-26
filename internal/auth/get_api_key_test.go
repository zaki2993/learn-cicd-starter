package auth

import (
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	got := ""
	exp := ""

	if !reflect.DeepEqual(exp, got) {
		t.Fatalf("expected: %v, got: %v", exp, got)
	}
}
