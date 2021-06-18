package problem1725

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
					{5, 8},
					{3, 9},
					{5, 12},
					{16, 5},
				},
			},
			output: output{
				value: 3,
			},
		},
		{
			input: input{
				value: [][]int{
					{2, 3},
					{3, 7},
					{4, 3},
					{3, 7},
				},
			},
			output: output{
				value: 3,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, countGoodRectangles(input.value))
	}
}
