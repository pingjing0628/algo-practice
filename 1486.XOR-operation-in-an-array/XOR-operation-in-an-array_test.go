package problem1486

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type example struct {
	input input
	output output
}

type input struct {
	value int
	start int
}

type output struct {
	value int
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: 5,
				start: 0,
			},
			output: output{
				value: 8,
			},
		},
		{
			input: input{
				value: 4,
				start: 3,
			},
			output: output{
				value: 8,
			},
		},
		{
			input: input{
				value: 1,
				start: 7,
			},
			output: output{
				value: 7,
			},
		},
		{
			input: input{
				value: 10,
				start: 5,
			},
			output: output{
				value: 2,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, xorOperation(input.value, input.start))
	}
}



