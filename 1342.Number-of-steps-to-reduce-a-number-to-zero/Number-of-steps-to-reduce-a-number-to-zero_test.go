package problem1342

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type example struct {
	input input
	output output
}

type input struct {
	value int
}

type output struct {
	value int
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: 14,
			},
			output: output{
				value: 6,
			},
		},
		{
			input: input{
				value: 8,
			},
			output: output{
				value: 4,
			},
		},
		{
			input: input{
				value: 123,
			},
			output: output{
				value: 12,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, numberOfSteps(input.value))
	}
}
