package problem1431

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
	extra int
}

type output struct {
	value []bool
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: []int{2, 3, 5, 1, 3},
				extra: 3,
			},
			output: output{
				value: []bool{true, true, true, false, true},
			},
		},
		{
			input: input{
				value: []int{4, 2, 1, 1, 2},
				extra: 1,
			},
			output: output{
				value: []bool{true, false, false, false, false},
			},
		},
		{
			input: input{
				value: []int{12, 1, 12},
				extra: 10,
			},
			output: output{
				value: []bool{true, false, true},
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, kidsWithCandies(input.value, input.extra))
	}
}


