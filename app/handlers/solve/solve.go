package solve

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Solve handles
func Solve(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "pong",
	})
}

type expressionResult struct {
	coefficient int
	constant    int
}

// Solve a linear equation with one unknown, x
// Rounds to 8 decimal places
func solveEquation(eqn string) (string, error) {
	// Remove spaces.
	eqn = strings.Replace(eqn, " ", "", -1)

	expressions := strings.Split(eqn, "=")

	if len(expressions) != 2 {
		return "", errors.New("Invalid equation")
	}

	left, err := evaluateExpression(expressions[0])
	right, err2 := evaluateExpression(expressions[1])

	if err != nil || err2 != nil {
		return "", errors.New("Invalid equation")
	}

	left.coefficient = left.coefficient - right.coefficient
	left.constant = right.constant - left.constant

	if left.coefficient == 0 && left.constant == 0 {
		return "Infinite solutions", nil
	}

	if left.coefficient == 0 {
		return "No solution", nil
	}

	result := fmt.Sprintf("%.8f", (float64(left.constant) / float64(left.coefficient)))
	// Trim trailing zeroes and decimal
	result = strings.TrimRight(strings.TrimRight(result, "0"), ".")

	return "x=" + result, nil
}

func evaluateExpression(exp string) (expressionResult, error) {
	tokens := strings.Split(strings.Replace(strings.Replace(exp, "-", "#-", -1), "+", "#+", -1), "#")
	coefficient := 0
	sum := 0

	if len(tokens) == 1 && tokens[0] == "" {
		return expressionResult{}, errors.New("Invalid expression")
	}

	for _, v := range tokens {
		if v == "" {
			continue
		} else if v == "+x" || v == "x" {
			coefficient++
		} else if v == "-x" {
			coefficient--
		} else if v[len(v)-1] == 'x' {
			i, err := strconv.ParseInt(v[:strings.Index(v, "x")], 10, 32)

			if err != nil {
				return expressionResult{}, errors.New("Invalid expression")
			}

			coefficient += int(i)
		} else {
			i, err := strconv.ParseInt(v, 10, 64)

			if err != nil {
				return expressionResult{}, errors.New("Invalid expression")
			}

			sum += int(i)
		}
	}
	return expressionResult{coefficient, sum}, nil
}
