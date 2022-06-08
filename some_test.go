package main

import (
	"os"
	"testing"
)

func TestSomething(t *testing.T) {
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
