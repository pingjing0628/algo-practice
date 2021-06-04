package problem1672

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
					{3, 2, 1},
				},
			},
			output: output{
				value: 6,
			},
		},
		{
			input: input{
				value: [][]int{
					{1, 5},
					{7, 3},
					{3, 5},
				},
			},
			output: output{
				value: 10,
			},
		},
		{
			input: input{
				value: [][]int{
					{2, 8, 7},
					{7, 1, 3},
					{1, 9, 5},
				},
			},
			output: output{
				value: 17,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, maximumWealth(input.value))
	}
}

