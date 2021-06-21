package problem0728

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type example struct {
	input input
	output output
}

type input struct {
	left int
	right int
}

type output struct {
	value []int
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				left: 1,
				right: 22,
			},
			output: output{
				value: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22},
			},
		},
		{
			input: input{
				left: 47,
				right: 85,
			},
			output: output{
				value: []int{48, 55, 66, 77},
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, selfDividingNumbers(input.left, input.right))
	}
}


