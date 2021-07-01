package problem1337

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
	k int
}

type output struct {
	value []int
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: [][]int{
					{1, 1, 0, 0, 0},
					{1, 1, 1, 1, 0},
					{1, 0, 0, 0, 0},
					{1, 1, 0, 0, 0},
					{1, 1, 1, 1, 1},
				},
				k: 3,
			},
			output: output{
				value: []int{2, 0, 3},
			},
		},
		{
			input: input{
				value: [][]int{
					{1, 0, 0, 0},
					{1, 1, 1, 1},
					{1, 0, 0, 0},
					{1, 0, 0, 0},
				},
				k: 2,
			},
			output: output{
				value: []int{0, 2},
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, kWeakestRows(input.value, input.k))
	}
}


