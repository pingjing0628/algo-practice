package problem1313

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
				value: []int{1, 2, 3, 4},
			},
			output: output{
				value: []int{2, 4, 4, 4},
			},
		},
		{
			input: input{
				value: []int{1, 1, 2, 3},
			},
			output: output{
				value: []int{1, 3, 3},
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, decompressRLElist(input.value))
	}
}

