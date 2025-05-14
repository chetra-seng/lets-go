package assert

import (
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func StringContains(t *testing.T, actual, expectedSubStr string) {
	t.Helper()

	if !strings.Contains(actual, expectedSubStr) {
		t.Errorf("got %q, want %q", actual, expectedSubStr)
	}
}

func NilError(t *testing.T, actual error) {
	t.Helper()

	if actual != nil {
		t.Errorf("got: %v; expected: nil", actual)
	}
}
