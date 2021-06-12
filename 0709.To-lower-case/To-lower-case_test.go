package problem0709

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type example struct {
	input input
	output output
}

type input struct {
	value string
}

type output struct {
	value string
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: "Hello",
			},
			output: output{
				value: "hello",
			},
		},
		{
			input: input{
				value: "here",
			},
			output: output{
				value: "here",
			},
		},
		{
			input: input{
				value: "LOVELY",
			},
			output: output{
				value: "lovely",
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, toLowerCase(input.value))
	}
}
