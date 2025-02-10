package main

import "testing"

func TestOddEve(t *testing.T) {
	if isEven(2) == false {
		t.Errorf("Expected %v to be even. Received false.", 2)
	}
}
