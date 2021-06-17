package problem1704

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
				value: "book",
			},
			output: output{
				value: true,
			},
		},
		{
			input: input{
				value: "textbook",
			},
			output: output{
				value: false,
			},
		},
		{
			input: input{
				value: "MerryChristmas",
			},
			output: output{
				value: false,
			},
		},
		{
			input: input{
				value: "AbCdEfGh",
			},
			output: output{
				value: true,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, halvesAreAlike(input.value))
	}
}

