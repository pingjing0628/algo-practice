package problem0961

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
				value: []int{1, 2, 3, 3},
			},
			output: output{
				value: 3,
			},
		},
		{
			input: input{
				value: []int{2, 1, 2, 5, 3, 2},
			},
			output: output{
				value: 2,
			},
		},
		{
			input: input{
				value: []int{5, 1, 5, 2, 5, 3, 5, 4},
			},
			output: output{
				value: 5,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, repeatedNTimes(input.value))
	}
}


