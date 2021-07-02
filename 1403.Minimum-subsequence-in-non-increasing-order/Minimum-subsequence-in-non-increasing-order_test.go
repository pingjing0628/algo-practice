package problem1403

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

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: []int{4, 3, 10, 9, 8},
			},
			output: output{
				value: []int{10, 9},
			},
		},
		{
			input: input{
				value: []int{4, 4, 7, 6, 7},
			},
			output: output{
				value: []int{7, 7, 6},
			},
		},
		{
			input: input{
				value: []int{6},
			},
			output: output{
				value: []int{6},
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, minSubsequence(input.value))
	}
}


