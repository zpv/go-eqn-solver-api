package solve

import (
	"fmt"
	"testing"
)

func TestSolveSimpleEquations(t *testing.T) {
	testCases := make(map[string]string)

	testCases["2x=4"] = "x=2"
	testCases["x=2"] = "x=2"

	for k, v := range testCases {
		result, err := solveEquation(k)

		if err != nil {
			t.Error("Error was not expected")
			return
		}

		assertEqual(t, result, v, "")
	}
}

func TestSolveComplexEquations(t *testing.T) {
	testCases := make(map[string]string)

	testCases["2x=4x+3x+x+1"] = "x=-0.16666667"
	testCases["5x+2x+1=2x+9"] = "x=1.6"

	for k, v := range testCases {
		result, err := solveEquation(k)

		if err != nil {
			t.Error("Error was not expected")
			return
		}

		assertEqual(t, result, v, "")
	}
}

func TestSolveEquationsWithSpaces(t *testing.T) {
	testCases := make(map[string]string)

	testCases["2x = 4x + 3x + x + 1"] = "x=-0.16666667"
	testCases["5x + 2x + 1 = 2x + 9"] = "x=1.6"

	for k, v := range testCases {
		result, err := solveEquation(k)

		if err != nil {
			t.Error("Error was not expected")
			return
		}

		assertEqual(t, result, v, "")
	}
}
func TestSolveEquationsWithDecimals(t *testing.T) {
	testCases := make(map[string]string)

	testCases["5x=4"] = "x=0.8"
	testCases["4x=8"] = "x=0.5"

	for k, v := range testCases {
		result, err := solveEquation(k)

		if err != nil {
			t.Error("Error was not expected")
			return
		}

		assertEqual(t, result, v, "")
	}
}

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}
