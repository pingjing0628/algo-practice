package problem1021

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
				value: "(()())(())",
			},
			output: output{
				value: "()()()",
			},
		},
		{
			input: input{
				value: "(()())(())(()(()))",
			},
			output: output{
				value: "()()()()(())",
			},
		},
		{
			input: input{
				value: "()()",
			},
			output: output{
				value: "",
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, removeOuterParentheses(input.value))
	}
}

