package main


import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if resut != "even" {
		t.Errorf("expected: even, actual: %s", result)
	}
}