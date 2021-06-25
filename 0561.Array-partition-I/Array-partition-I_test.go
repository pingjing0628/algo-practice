package problem0561

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
				value: []int{1, 4, 3, 2},
			},
			output: output{
				value: 4,
			},
		},
		{
			input: input{
				value: []int{6, 2, 6, 5, 1, 2},
			},
			output: output{
				value: 9,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, arrayPairSum(input.value))
	}
}


