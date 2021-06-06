package problem1528

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
	indices []int
}

type output struct {
	value string
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: "codeleet",
				indices: []int{4, 5, 6, 7, 0, 2, 1, 3},
			},
			output: output{
				value: "leetcode",
			},
		},
		{
			input: input{
				value: "aiohn",
				indices: []int{3, 1, 4, 2, 0},
			},
			output: output{
				value: "nihao",
			},
		},
		{
			input: input{
				value: "aaiougrt",
				indices: []int{4, 0, 2, 6, 7, 3, 1, 5},
			},
			output: output{
				value: "arigatou",
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, restoreString(input.value, input.indices))
	}
}

