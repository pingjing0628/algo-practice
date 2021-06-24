package problem1299

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
				value: []int{17, 18, 5, 4, 6, 1},
			},
			output: output{
				value: []int{18, 6, 6, 6, 1, -1},
			},
		},
		{
			input: input{
				value: []int{400},
			},
			output: output{
				value: []int{-1},
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, replaceElements(input.value))
	}
}


