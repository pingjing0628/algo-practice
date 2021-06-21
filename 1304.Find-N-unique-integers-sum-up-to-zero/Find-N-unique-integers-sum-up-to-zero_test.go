package problem1304

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type example struct {
	input input
	output output
}

type input struct {
	value int
}

type output struct {
	value []int
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: 5,
			},
			output: output{
				value: []int{1, -1, 2, -2, 0},
			},
		},
		{
			input: input{
				value: 3,
			},
			output: output{
				value: []int{1, -1, 0},
			},
		},
		{
			input: input{
				value: 1,
			},
			output: output{
				value: []int{0},
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, sumZero(input.value))
	}
}

