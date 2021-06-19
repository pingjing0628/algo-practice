package problem1572

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

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			output: output{
				value: 25,
			},
		},
		{
			input: input{
				value: [][]int{
					{1, 1, 1, 1},
					{1, 1, 1, 1},
					{1, 1, 1, 1},
					{1, 1, 1, 1},
				},
			},
			output: output{
				value: 8,
			},
		},
		{
			input: input{
				value: [][]int{
					{5},
				},
			},
			output: output{
				value: 5,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, diagonalSum(input.value))
	}
}

