package problem1047

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
				value: "abbaca",
			},
			output: output{
				value: "ca",
			},
		},
		{
			input: input{
				value: "azxxzy",
			},
			output: output{
				value: "ay",
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, removeDuplicates(input.value))
	}
}


