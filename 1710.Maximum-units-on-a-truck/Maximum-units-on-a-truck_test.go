package problem1710

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type example struct {
	input input
	output output
}

type input struct {
	value [][]int
	truckSize int
}

type output struct {
	value int
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: [][]int{
					{1, 3},
					{2, 2},
					{3, 1},
				},
				truckSize: 4,
			},
			output: output{
				value: 8,
			},
		},
		{
			input: input{
				value: [][]int{
					{5, 10},
					{2, 5},
					{4, 7},
					{3, 9},
				},
				truckSize: 10,
			},
			output: output{
				value: 91,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, maximumUnits(input.value, input.truckSize))
	}
}

