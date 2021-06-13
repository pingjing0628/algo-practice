package problem1844

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
				value: "a1c1e1",
			},
			output: output{
				value: "abcdef",
			},
		},
		{
			input: input{
				value: "a1b2c3d4e",
			},
			output: output{
				value: "abbdcfdhe",
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, replaceDigits(input.value))
	}
}

