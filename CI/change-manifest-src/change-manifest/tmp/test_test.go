package main

import (
	"testing"
)


func TestHelloName(t *testing.T) {
	if 2 != 1 {
		t.Errorf("Result was incorrect, got: %s, want: %s.", "1", "Foo")
	}
}
