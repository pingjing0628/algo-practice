package problem1470

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
	n int
}

type output struct {
	value []int
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: []int{2, 5, 1, 3, 4, 7},
				n: 3,
			},
			output: output{
				value: []int{2, 3, 5, 4, 1, 7},
			},
		},
		{
			input: input{
				value: []int{1, 2, 3, 4, 4, 3, 2, 1},
				n: 4,
			},
			output: output{
				value: []int{1, 4, 2, 3, 3, 2, 4, 1},
			},
		},
		{
			input: input{
				value: []int{1, 1, 2, 2},
				n: 2,
			},
			output: output{
				value: []int{1, 2, 1, 2},
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, shuffle(input.value, input.n))
	}
}

