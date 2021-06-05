package problem1512

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
				value: []int{1, 2, 3, 1, 1, 3},
			},
			output: output{
				value: 4,
			},
		},
		{
			input: input{
				value: []int{1, 1, 1, 1},
			},
			output: output{
				value: 6,
			},
		},
		{
			input: input{
				value: []int{1, 2, 3},
			},
			output: output{
				value: 0,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, numIdenticalPairs(input.value))
	}
}

