package problem1614

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
				value: "(1+(2*3)+((8)/4))+1",
			},
			output: output{
				value: 3,
			},
		},
		{
			input: input{
				value: "(1)+((2))+(((3)))",
			},
			output: output{
				value: 3,
			},
		},
		{
			input: input{
				value: "1+(2*3)/(2-1)",
			},
			output: output{
				value: 1,
			},
		},
		{
			input: input{
				value: "1",
			},
			output: output{
				value: 0,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, maxDepth(input.value))
	}
}


