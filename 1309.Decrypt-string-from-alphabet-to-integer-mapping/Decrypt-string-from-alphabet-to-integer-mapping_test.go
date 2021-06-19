package problem1309

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
	value string
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: "10#11#12",
			},
			output: output{
				value: "jkab",
			},
		},
		{
			input: input{
				value: "1326#",
			},
			output: output{
				value: "acz",
			},
		},
		{
			input: input{
				value: "25#",
			},
			output: output{
				value: "y",
			},
		},
		{
			input: input{
				value: "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#",
			},
			output: output{
				value: "abcdefghijklmnopqrstuvwxyz",
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, freqAlphabets(input.value))
	}
}


