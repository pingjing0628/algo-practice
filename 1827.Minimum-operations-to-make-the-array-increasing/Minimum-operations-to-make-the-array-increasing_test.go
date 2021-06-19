package problem1827

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type example struct {
	input input
	output output
}

type input struct {
	value []int
}

type output struct {
	value int
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: []int{1, 1, 1},
			},
			output: output{
				value: 3,
			},
		},
		{
			input: input{
				value: []int{1, 5, 2, 4, 1},
			},
			output: output{
				value: 14,
			},
		},
		{
			input: input{
				value: []int{8},
			},
			output: output{
				value: 0,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, minOperations(input.value))
	}
}


