package problem0657

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
				value: "UD",
			},
			output: output{
				value: true,
			},
		},
		{
			input: input{
				value: "LL",
			},
			output: output{
				value: false,
			},
		},
		{
			input: input{
				value: "RRDD",
			},
			output: output{
				value: false,
			},
		},
		{
			input: input{
				value: "LDRRLRUULR",
			},
			output: output{
				value: false,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, judgeCircle(input.value))
	}
}

