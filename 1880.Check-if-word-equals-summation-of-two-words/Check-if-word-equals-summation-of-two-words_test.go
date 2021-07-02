package problem1880

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type example struct {
	input input
	output output
}

type input struct {
	firstWord string
	secondWord string
	targetWord string
}

type output struct {
	value bool
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				firstWord: "acb",
				secondWord: "cba",
				targetWord: "cdb",
			},
			output: output{
				value: true,
			},
		},
		{
			input: input{
				firstWord: "aaa",
				secondWord: "a",
				targetWord: "aab",
			},
			output: output{
				value: false,
			},
		},
		{
			input: input{
				firstWord: "aaa",
				secondWord: "a",
				targetWord: "aaaa",
			},
			output: output{
				value: true,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, isSumEqual(input.firstWord, input.secondWord, input.targetWord))
	}
}

