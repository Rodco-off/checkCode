package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTestCase(t *testing.T) {
	task := Task{TestCases: []byte(`[{"input": "2 3", "output": "5"}]`)}

	result, err := task.ParseTestCases()

	assert.NoError(t, err)

	assert.Len(t, result, 1)
	assert.Equal(t, "2 3", result[0].Input)
	assert.Equal(t, "5", result[0].Output)

	badTask := Task{TestCases: []byte("Невалидный json")}

	_, err = badTask.ParseTestCases()
	assert.Error(t, err)
}
