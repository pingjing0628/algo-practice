package problem1475

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
	value []int
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: []int{8, 4, 6, 2, 3},
			},
			output: output{
				value: []int{4, 2, 4, 2, 3},
			},
		},
		{
			input: input{
				value: []int{1, 2, 3, 4, 5},
			},
			output: output{
				value: []int{1, 2, 3, 4, 5},
			},
		},
		{
			input: input{
				value: []int{10, 1, 1, 6},
			},
			output: output{
				value: []int{9, 0, 1, 6},
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, finalPrices(input.value))
	}
}

