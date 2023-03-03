package main

import "testing"

func TestMySum(t *testing.T) {
	x := mySum(2, 3)
	if x != 5 {
		t.Errorf("Expected = %d; want 5", x)
		// t.Error("Expected 5, got %v", x)
	}
}
