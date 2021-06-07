package problem1720

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
	first int
}

type output struct {
	value []int
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: []int{1, 2, 3},
				first: 1,
			},
			output: output{
				value: []int{1, 0, 2, 1},
			},
		},
		{
			input: input{
				value: []int{6, 2, 7, 3},
				first: 4,
			},
			output: output{
				value: []int{4, 2, 0, 7, 4},
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, decode(input.value, input.first))
	}
}

