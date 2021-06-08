package problem1773

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
	ruleKey string
	ruleValue string
}

type output struct {
	value int
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: [][]string{
					{"phone", "blue", "pixel"},
					{"computer", "silver", "lenovo"},
					{"phone", "gold", "iphone"},
				},
				ruleKey: "color",
				ruleValue: "silver",
			},
			output: output{
				value: 1,
			},
		},
		{
			input: input{
				value: [][]string{
					{"phone", "blue", "pixel"},
					{"computer", "silver", "lenovo"},
					{"phone", "gold", "iphone"},
				},
				ruleKey: "type",
				ruleValue: "phone",
			},
			output: output{
				value: 2,
			},
		},
		{
			input: input{
				value: [][]string{
					{"phone", "blue", "pixel"},
					{"computer", "silver", "lenovo"},
					{"phone", "gold", "iphone"},
				},
				ruleKey: "name",
				ruleValue: "ipad",
			},
			output: output{
				value: 0,
			},
		},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, countMatches(input.value, input.ruleKey, input.ruleValue))
	}
}


