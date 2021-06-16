package problem0804

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type example struct {
	input input
	output output
}

type input struct {
	value []string
}

type output struct {
	value int
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: []string{
					"gin",
					"zen",
					"gig",
					"msg",
				},
			},
			output: output{
				value: 2,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, uniqueMorseRepresentations(input.value))
	}
}

