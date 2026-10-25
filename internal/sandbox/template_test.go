package sandbox

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerativeMainFile(t *testing.T) {
	userCode := `func Solve(input string) string {return "hello"}`

	result, err := GenerativeMainFile(userCode)

	assert.NoError(t, err)

	assert.True(t, strings.Contains(result, "package main"))
	assert.True(t, strings.Contains(result, userCode))
	assert.True(t, strings.Contains(result, "func main()"))
}
