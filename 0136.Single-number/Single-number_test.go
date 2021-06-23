package problem0136

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
				value: []int{2, 2, 1},
			},
			output: output{
				value: 1,
			},
		},
		{
			input: input{
				value: []int{4, 1, 2, 1, 2},
			},
			output: output{
				value: 4,
			},
		},
		{
			input: input{
				value: []int{1},
			},
			output: output{
				value: 1,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, singleNumber(input.value))
	}
}
