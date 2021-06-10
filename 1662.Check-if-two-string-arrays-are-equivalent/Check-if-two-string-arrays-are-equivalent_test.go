package problem1662

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type example struct {
	input input
	output output
}

type input struct {
	word1 []string
	word2 []string
}

type output struct {
	value bool
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				word1: []string{
					"ab",
					"c",
				},
				word2: []string{
					"a",
					"bc",
				},
			},
			output: output{
				value: true,
			},
		},
		{
			input: input{
				word1: []string{
					"ac",
					"b",
				},
				word2: []string{
					"ab",
					"c",
				},
			},
			output: output{
				value: false,
			},
		},
		{
			input: input{
				word1: []string{
					"abc",
					"d",
					"defg",
				},
				word2: []string{
					"abcddefg",
				},
			},
			output: output{
				value: true,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, arrayStringsAreEqual(input.word1, input.word2))
	}
}


