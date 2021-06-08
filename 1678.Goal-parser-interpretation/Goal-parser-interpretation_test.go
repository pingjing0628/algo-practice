package problem1678

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
				value: "G()(al)",
			},
			output: output{
				value: "Goal",
			},
		},
		{
			input: input{
				value: "G()()()()(al)",
			},
			output: output{
				value: "Gooooal",
			},
		},
		{
			input: input{
				value: "(al)G(al)()()G",
			},
			output: output{
				value: "alGalooG",
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, interpret(input.value))
	}
}


