package problem0832

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type example struct {
	input input
	output output
}

type input struct {
	value [][]int
}

type output struct {
	value [][]int
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: [][]int{
					{1, 1, 0},
					{1, 0, 1},
					{0, 0, 0},
				},
			},
			output: output{
				value: [][]int{
					{1, 0, 0},
					{0, 1, 0},
					{1, 1, 1},
				},
			},
		},
		{
			input: input{
				value: [][]int{
					{1, 1, 0, 0},
					{1, 0, 0, 1},
					{0, 1, 1, 1},
					{1, 0, 1, 0},
				},
			},
			output: output{
				value: [][]int{
					{1, 1, 0, 0},
					{0, 1, 1, 0},
					{0, 0, 0, 1},
					{1, 0, 1, 0},
				},
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, flipAndInvertImage(input.value))
	}
}
