package problem1742

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type example struct {
	input input
	output output
}

type input struct {
	value1 int
	value2 int
}

type output struct {
	value int
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value1: 1,
				value2: 10,
			},
			output: output{
				value: 2,
			},
		},
		{
			input: input{
				value1: 5,
				value2: 15,
			},
			output: output{
				value: 2,
			},
		},
		{
			input: input{
				value1: 19,
				value2: 28,
			},
			output: output{
				value: 2,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, countBalls(input.value1, input.value2))
	}
}
