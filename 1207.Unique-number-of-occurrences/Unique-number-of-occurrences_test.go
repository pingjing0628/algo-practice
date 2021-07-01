package problem1207

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
	value bool
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: []int{1, 2, 2, 1, 1, 3},
			},
			output: output{
				value: true,
			},
		},
		{
			input: input{
				value: []int{1, 2},
			},
			output: output{
				value: false,
			},
		},
		{
			input: input{
				value: []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			},
			output: output{
				value: true,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, uniqueOccurrences(input.value))
	}
}



