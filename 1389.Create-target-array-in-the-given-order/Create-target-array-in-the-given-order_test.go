package problem1389

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
	index []int
}

type output struct {
	value []int
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: []int{0, 1, 2, 3, 4},
				index: []int{0, 1, 2, 2, 1},
			},
			output: output{
				value: []int{0, 4, 1, 3, 2},
			},
		},
		{
			input: input{
				value: []int{1, 2, 3, 4, 0},
				index: []int{0, 1, 2, 3, 0},
			},
			output: output{
				value: []int{0, 1, 2, 3, 4},
			},
		},
		{
			input: input{
				value: []int{1},
				index: []int{0},
			},
			output: output{
				value: []int{1},
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, createTargetArray(input.value, input.index))
	}
}

