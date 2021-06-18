package problem1295

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
				value: []int{12, 345, 2, 6, 7896},
			},
			output: output{
				value: 2,
			},
		},
		{
			input: input{
				value: []int{555, 901, 482, 1771},
			},
			output: output{
				value: 1,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, findNumbers(input.value))
	}
}
