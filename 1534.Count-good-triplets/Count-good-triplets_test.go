package problem1534

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type example struct {
	input input
	output output
}

type input struct {
	value []int
	a int
	b int
	c int
}

type output struct {
	value int
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: []int{3,0,1,1,9,7},
				a: 7,
				b: 2,
				c: 3,
			},
			output: output{
				value: 4,
			},
		},
		{
			input: input{
				value: []int{1,1,2,2,3},
				a: 0,
				b: 0,
				c: 1,
			},
			output: output{
				value: 0,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, countGoodTriplets(input.value, input.a, input.b, input.c))
	}
}
