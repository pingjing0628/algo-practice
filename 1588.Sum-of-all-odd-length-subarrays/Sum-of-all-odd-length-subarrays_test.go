package problem1588

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
				value: []int{1, 4, 2, 5, 3},
			},
			output: output{
				value: 58,
			},
		},
		{
			input: input{
				value: []int{1, 2},
			},
			output: output{
				value: 3,
			},
		},
		{
			input: input{
				value: []int{10, 11, 12},
			},
			output: output{
				value: 66,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, sumOddLengthSubarrays(input.value))
	}
}



