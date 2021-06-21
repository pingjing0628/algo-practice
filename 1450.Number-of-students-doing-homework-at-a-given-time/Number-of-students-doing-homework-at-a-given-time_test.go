package problem1450

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type example struct {
	input input
	output output
}

type input struct {
	startTime []int
	endTime []int
	queryTime int
}

type output struct {
	value int
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				startTime: []int{1, 2, 3},
				endTime: []int{3, 2, 7},
				queryTime: 4,
			},
			output: output{
				value: 1,
			},
		},
		{
			input: input{
				startTime: []int{4},
				endTime: []int{4},
				queryTime: 4,
			},
			output: output{
				value: 1,
			},
		},
		{
			input: input{
				startTime: []int{4},
				endTime: []int{4},
				queryTime: 5,
			},
			output: output{
				value: 0,
			},
		},
		{
			input: input{
				startTime: []int{1, 1, 1, 1},
				endTime: []int{1, 3, 2, 4},
				queryTime: 7,
			},
			output: output{
				value: 0,
			},
		},
		{
			input: input{
				startTime: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
				endTime: []int{10, 10, 10, 10, 10, 10, 10, 10, 10},
				queryTime: 5,
			},
			output: output{
				value: 5,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, busyStudent(input.startTime, input.endTime, input.queryTime))
	}
}

