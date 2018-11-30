package solve

import (
	"errors"
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
	constants   int
}

// Solve a linear equation with one unknown, x
func solveEquation(eqn string) (string, error) {
	// Remove spaces.
	eqn = strings.Replace(eqn, " ", "", -1)

	if !validString(eqn) {
		return "", errors.New("eqn: invalid linear equation")
	}

	return "", nil
}

func evaluateExpression(exp string) expressionResult {
	// stub
	return expressionResult{0, 0}
}

func validString(eqn string) bool {
	// stub
	return false
}
