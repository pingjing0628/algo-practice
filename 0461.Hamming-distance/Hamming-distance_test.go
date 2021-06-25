package problem0461

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type example struct {
	input input
	output output
}

type input struct {
	x int
	y int
}

type output struct {
	value int
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				x: 1,
				y: 4,
			},
			output: output{
				value: 2,
			},
		},
		{
			input: input{
				x: 3,
				y: 1,
			},
			output: output{
				value: 1,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, hammingDistance(input.x, input.y))
	}
}



