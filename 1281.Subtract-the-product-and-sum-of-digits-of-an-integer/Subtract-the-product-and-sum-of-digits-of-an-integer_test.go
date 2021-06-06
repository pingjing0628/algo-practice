package problem1281

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
				value: 234,
			},
			output: output{
				value: 15,
			},
		},
		{
			input: input{
				value: 4421,
			},
			output: output{
				value: 21,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, subtractProductAndSum(input.value))
	}
}

