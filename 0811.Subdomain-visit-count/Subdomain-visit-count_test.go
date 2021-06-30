package problem0811

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type example struct {
	input input
	output output
}

type input struct {
	value []string
}

type output struct {
	value []string
}

func TestSingle(t *testing.T)  {
	examples := []example {
		{
			input: input{
				value: []string{"9001 discuss.leetcode.com"},
			},
			output: output{
				value: []string{"9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"},
			},
		},
		//{
		//	input: input{
		//		value: []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"},
		//	},
		//	output: output{
		//		value: []string{
		//			"900 google.mail.com",
		//			"901 mail.com",
		//			"951 com",
		//			"50 yahoo.com",
		//			"1 intel.mail.com",
		//			"5 wiki.org",
		//			"5 org",
		//		},
		//	},
		//},
	}

	for _, val := range examples{
		output, input := val.output, val.input
		assert.Equal(t, output.value, subdomainVisits(input.value))
	}
}


