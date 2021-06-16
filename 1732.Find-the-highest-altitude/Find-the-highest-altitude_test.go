package problem1732

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
				value: []int{-5, 1, 5, 0, -7},
			},
			output: output{
				value: 1,
			},
		},
		{
			input: input{
				value: []int{-4, -3, -2, -1, 4, 3, 2},
			},
			output: output{
				value: 0,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, largestAltitude(input.value))
	}
}
