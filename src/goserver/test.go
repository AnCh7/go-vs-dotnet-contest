package main

import "testing"

func TestMethod(t *testing.T) {
	expected := true
	actual := true
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}
