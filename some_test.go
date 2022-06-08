package main

import (
	"os"
	"testing"
)

func TestLintShouldFailButDoesnt(t *testing.T) {
	tests := map[string]struct {
		fooValue string
	}{
		"does something": {fooValue: "bar"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			os.Setenv("foo", tc.fooValue)
		})
	}
}
func TestLintFails(t *testing.T) {
	os.Setenv("foo", "bar")
}
