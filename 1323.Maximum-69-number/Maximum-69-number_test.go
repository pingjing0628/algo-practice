package problem1323

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
}

type output struct {
	value int
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: 9669,
			},
			output: output{
				value: 9969,
			},
		},
		{
			input: input{
				value: 9996,
			},
			output: output{
				value: 9999,
			},
		},
		{
			input: input{
				value: 9999,
			},
			output: output{
				value: 9999,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, maximum69Number(input.value))
	}
}

