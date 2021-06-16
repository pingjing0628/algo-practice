package problem1266

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
					{1, 1},
					{3, 4},
					{-1, 0},
				},
			},
			output: output{
				value: 7,
			},
		},
		{
			input: input{
				value: [][]int{
					{3, 2},
					{-2, 2},
				},
			},
			output: output{
				value: 5,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, minTimeToVisitAllPoints(input.value))
	}
}

