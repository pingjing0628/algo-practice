package problem1688

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
				value: 7,
			},
			output: output{
				value: 6,
			},
		},
		{
			input: input{
				value: 14,
			},
			output: output{
				value: 13,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, numberOfMatches(input.value))
	}
}



