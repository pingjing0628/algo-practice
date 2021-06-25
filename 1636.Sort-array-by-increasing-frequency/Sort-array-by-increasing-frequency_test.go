package problem1636

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

func TestSort(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: []int{1, 1, 2, 2, 2, 3},
			},
			output: output{
				value: []int{3, 1, 1, 2, 2, 2},
			},
		},
		{
			input: input{
				value: []int{2, 3, 1, 3, 2},
			},
			output: output{
				value: []int{1, 3, 3, 2, 2},
			},
		},
		{
			input: input{
				value: []int{-1, 1, -6, 4, 5, -6, 1, 4, 1},
			},
			output: output{
				value: []int{5, -1, 4, 4, -6, -6, 1, 1, 1},
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, frequencySort(input.value))
	}
}
