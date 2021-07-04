package problem0922

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
				value: []int{3, 1, 2, 4},
			},
			output: output{
				value: []int{4, 3, 2, 1},
			},
		},
		{
			input: input{
				value: []int{4, 2, 5, 7},
			},
			output: output{
				value: []int{4, 5, 2, 7},
			},
		},
		{
			input: input{
				value: []int{2, 3},
			},
			output: output{
				value: []int{2, 3},
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, sortArrayByParityII(input.value))
	}
}

