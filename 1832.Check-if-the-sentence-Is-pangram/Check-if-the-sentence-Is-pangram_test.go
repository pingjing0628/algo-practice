package problem1832

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type example struct {
	input input
	output output
}

type input struct {
	value string
}

type output struct {
	value bool
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: "thequickbrownfoxjumpsoverthelazydog",
			},
			output: output{
				value: true,
			},
		},
		{
			input: input{
				value: "leetcode",
			},
			output: output{
				value: false,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, checkIfPangram(input.value))
	}
}

