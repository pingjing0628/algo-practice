package problem1464

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
				value: []int{3, 4, 5, 2},
			},
			output: output{
				value: 12,
			},
		},
		{
			input: input{
				value: []int{1, 5, 4, 5},
			},
			output: output{
				value: 16,
			},
		},
		{
			input: input{
				value: []int{3, 7},
			},
			output: output{
				value: 12,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, maxProduct(input.value))
	}
}

