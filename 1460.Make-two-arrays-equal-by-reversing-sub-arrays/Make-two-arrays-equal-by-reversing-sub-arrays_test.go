package problem1460

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
	arr []int
}

type output struct {
	value bool
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: []int{1, 2, 3, 4},
				arr: []int{2, 4, 1, 3},
			},
			output: output{
				value: true,
			},
		},
		{
			input: input{
				value: []int{1, 12},
				arr: []int{12, 1},
			},
			output: output{
				value: true,
			},
		},
		{
			input: input{
				value: []int{3, 7, 9},
				arr: []int{3, 7, 11},
			},
			output: output{
				value: false,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, canBeEqual(input.value, input.arr))
	}
}


