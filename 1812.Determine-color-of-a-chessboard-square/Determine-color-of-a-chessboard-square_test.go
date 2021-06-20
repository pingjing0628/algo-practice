package problem1812

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
	value bool
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: "a1",
			},
			output: output{
				value: false,
			},
		},
		{
			input: input{
				value: "h3",
			},
			output: output{
				value: true,
			},
		},
		{
			input: input{
				value: "c7",
			},
			output: output{
				value: false,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, squareIsWhite(input.value))
	}
}
