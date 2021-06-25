package problem1768

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type example struct {
	input input
	output output
}

type input struct {
	word1 string
	word2 string
}

type output struct {
	value string
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				word1: "abc",
				word2: "pqr",
			},
			output: output{
				value: "apbqcr",
			},
		},
		{
			input: input{
				word1: "ab",
				word2: "pqrs",
			},
			output: output{
				value: "apbqrs",
			},
		},
		{
			input: input{
				word1: "abcd",
				word2: "pq",
			},
			output: output{
				value: "apbqcd",
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, mergeAlternately(input.word1, input.word2))
	}
}


