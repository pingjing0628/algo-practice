package problem1684

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
	words []string
}

type output struct {
	value int
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: "ab",
				words: []string{
					"ad",
					"bd",
					"aaab",
					"baa",
					"badab",
				},
			},
			output: output{
				value: 2,
			},
		},
		{
			input: input{
				value: "abc",
				words: []string{
					"a",
					"b",
					"c",
					"ab",
					"ac",
					"bc",
					"abc",
				},
			},
			output: output{
				value: 7,
			},
		},
		{
			input: input{
				value: "cad",
				words: []string{
					"cc",
					"acd",
					"b",
					"ba",
					"bac",
					"bad",
					"ac",
					"d",
				},
			},
			output: output{
				value: 4,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, countConsistentStrings(input.value, input.words))
	}
}



