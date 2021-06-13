package problem1816

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
	k int
}

type output struct {
	value string
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: "Hello how are you Contestant",
				k: 4,
			},
			output: output{
				value: "Hello how are you",
			},
		},
		{
			input: input{
				value: "What is the solution to this problem",
				k: 4,
			},
			output: output{
				value: "What is the solution",
			},
		},
		{
			input: input{
				value: "chopper is not a tanuki",
				k: 5,
			},
			output: output{
				value: "chopper is not a tanuki",
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, truncateSentence(input.value, input.k))
	}
}

