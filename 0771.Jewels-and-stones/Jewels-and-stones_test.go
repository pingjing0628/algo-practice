package problem0771

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type example struct {
	input input
	output output
}

type input struct {
	value1 string
	value2 string
}

type output struct {
	value int
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value1: "aA",
				value2: "aAAbbbb",
			},
			output: output{
				value: 3,
			},
		},
		{
			input: input{
				value1: "z",
				value2: "ZZ",
			},
			output: output{
				value: 0,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, numJewelsInStones(input.value1, input.value2))
	}
}
