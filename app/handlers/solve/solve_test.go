package solve

import "testing"

func TestSolveEquation(t *testing.T) {
	result, err := solveEquation("2x=4")

	if err != nil {
		t.Error("Error was not expected")
		return
	}

	if result != "x=2" {
		t.Error("Wrong result expected")
	}
}
