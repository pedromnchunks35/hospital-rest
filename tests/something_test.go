package tests

import "testing"

func TestAdd(t *testing.T) {
	x := 4 * 3

	expected := 12

	if x != expected {
		t.Errorf("Should be a given value")
	}
}
