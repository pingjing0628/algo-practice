package problem1374

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
	value string
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: 4,
			},
			output: output{
				value: "aaab",
			},
		},
		{
			input: input{
				value: 2,
			},
			output: output{
				value: "ab",
			},
		},
		{
			input: input{
				value: 7,
			},
			output: output{
				value: "aaaaaaa",
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, generateTheString(input.value))
	}
}

