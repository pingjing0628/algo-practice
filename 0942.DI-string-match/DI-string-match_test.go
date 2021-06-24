package problem0942

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
	value []int
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: "IDID",
			},
			output: output{
				value: []int{0, 4, 1, 3, 2},
			},
		},
		{
			input: input{
				value: "III",
			},
			output: output{
				value: []int{0, 1, 2, 3},
			},
		},
		{
			input: input{
				value: "DDI",
			},
			output: output{
				value: []int{3, 2, 0, 1},
			},
		},
		{
			input: input{
				value: "DIDDIDIDID",
			},
			output: output{
				value: []int{10, 0, 9, 8, 1, 7, 2, 6, 3, 5, 4},
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, diStringMatch(input.value))
	}
}


