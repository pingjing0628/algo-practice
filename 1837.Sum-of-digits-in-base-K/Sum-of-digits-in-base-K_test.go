package problem1837

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type example struct {
	input input
	output output
}

type input struct {
	n int
	k int
}

type output struct {
	value int
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				n: 34,
				k: 6,
			},
			output: output{
				value: 9,
			},
		},
		{
			input: input{
				n: 10,
				k: 10,
			},
			output: output{
				value: 1,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, sumBase(input.n, input.k))
	}
}

