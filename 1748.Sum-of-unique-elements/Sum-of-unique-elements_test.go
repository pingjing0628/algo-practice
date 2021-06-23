package problem1748

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
				value: []int{1, 2, 3, 2},
			},
			output: output{
				value: 4,
			},
		},
		{
			input: input{
				value: []int{1, 1, 1, 1, 1},
			},
			output: output{
				value: 0,
			},
		},
		{
			input: input{
				value: []int{1, 2, 3, 4, 5},
			},
			output: output{
				value: 15,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, sumOfUnique(input.value))
	}
}


