package problem1351

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type example struct {
	input input
	output output
}

type input struct {
	value [][]int
}

type output struct {
	value int
}

func TestNegative(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: [][]int{
					{4, 3, 2, -1},
					{3, 2, 1, -1},
					{1, 1, -1, -2},
					{-1, -1, -2, -3},
				},
			},
			output: output{
				value: 8,
			},
		},
		{
			input: input{
				value: [][]int{
					{3, 2},
					{1, 0},
				},
			},
			output: output{
				value: 0,
			},
		},
		{
			input: input{
				value: [][]int{
					{1, -1},
					{-1, -1},
				},
			},
			output: output{
				value: 3,
			},
		},
		{
			input: input{
				value: [][]int{
					{-1},
				},
			},
			output: output{
				value: 1,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, countNegatives(input.value))
	}
}