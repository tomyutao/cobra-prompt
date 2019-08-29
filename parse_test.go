package cobraprompt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test case
func TestParse(t *testing.T) {
	assert := assert.New(t)
	input := "-v --format \"some example\" -i test"
	output, err := parseCommandLine(input)
	expect := []string{"-v", "--format", "some example", "-i", "test"}
	if err != nil {
		t.Errorf("Err [%v] from %v", err, input)
	} else {
		assert.Equal(output, expect, "Two arrays should be the same")
	}
}
