package utils

import "testing"

func TestSquareRoot(t *testing.T) {
	expectedSquareRoot := 5.477225575051661
	actualSquareRoot := SquareRoot(30)

	if actualSquareRoot != expectedSquareRoot {
		t.Errorf("Invalid square root. Expected %g, got %g", expectedSquareRoot, actualSquareRoot)
	}
}

