package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		len  int
	}{
		{"", 0},
		{"zekai", 5},
		{"xinran", 6},
		{"yuwei", 5},
	}

	for _, tests := range tests {
		if actual := getNameLen(tests.name); actual != tests.len {
			t.Errorf("getNameLen(%s); "+"got %d, expected %d.", tests.name, actual, tests.len)
		}
	}
}
