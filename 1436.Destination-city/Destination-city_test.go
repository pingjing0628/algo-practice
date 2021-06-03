package problem1436

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type example struct {
	input input
	output output
}

type input struct {
	value [][]string
}

type output struct {
	value string
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: [][]string{
					{"pYyNGfBYbm", "wxAscRuzOl"},
					{"kzwEQHfwce", "pYyNGfBYbm"},
				},
			},
			output: output{
				value: "wxAscRuzOl",
			},
		},
		{
			input: input{
				value: [][]string{
					{"jMgaf WaWA", "iinynVdmBz"},
					{" QCrEFBcAw", "wRPRHznLWS"},
					{"iinynVdmBz", "OoLjlLFzjz"},
					{"OoLjlLFzjz", " QCrEFBcAw"},
					{"IhxjNbDeXk", "jMgaf WaWA"},
					{"jmuAYy vgz", "IhxjNbDeXk"},
				},
			},
			output: output{
				value: "wRPRHznLWS",
			},
		},
		{
			input: input{
				value: [][]string{
					{"London", "New York"},
					{"New York", "Lima"},
					{"Lima", "Sao Paulo"},
				},
			},
			output: output{
				value: "Sao Paulo",
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, destCity(input.value))
	}
}
