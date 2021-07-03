package problem0852

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
				value: []int{0, 1, 0},
			},
			output: output{
				value: 1,
			},
		},
		{
			input: input{
				value: []int{0, 2, 1, 0},
			},
			output: output{
				value: 1,
			},
		},
		{
			input: input{
				value: []int{0, 10, 5, 2},
			},
			output: output{
				value: 1,
			},
		},
		{
			input: input{
				value: []int{3, 4, 5, 1},
			},
			output: output{
				value: 2,
			},
		},
		{
			input: input{
				value: []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19},
			},
			output: output{
				value: 2,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, peakIndexInMountainArray(input.value))
	}
}


