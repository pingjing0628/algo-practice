package problem1370

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

func TestNegative(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: "aaaabbbbcccc",
			},
			output: output{
				value: "abccbaabccba",
			},
		},
		{
			input: input{
				value: "leetcode",
			},
			output: output{
				value: "cdelotee",
			},
		},
		{
			input: input{
				value: "ggggggg",
			},
			output: output{
				value: "ggggggg",
			},
		},
		{
			input: input{
				value: "spo",
			},
			output: output{
				value: "ops",
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, sortString(input.value))
	}
}