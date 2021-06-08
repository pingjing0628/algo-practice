package problem1221

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
	value int
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: "RLRRLLRLRL",
			},
			output: output{
				value: 4,
			},
		},
		{
			input: input{
				value: "RLLLLRRRLR",
			},
			output: output{
				value: 3,
			},
		},
		{
			input: input{
				value: "LLLLRRRR",
			},
			output: output{
				value: 1,
			},
		},
		{
			input: input{
				value: "RLRRRLLRLL",
			},
			output: output{
				value: 2,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, balancedStringSplit(input.value))
	}
}


