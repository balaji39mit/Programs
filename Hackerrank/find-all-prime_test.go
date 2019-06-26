package hackerrank

import (
	"testing"

	"fmt"
)

func TestCountPrime(t *testing.T) {
	testData := []struct {
		input  int64
		output int32
	}{
		{
			input:  500,
			output: 4,
		},
		{
			input:  1,
			output: 0,
		},
		{
			input:  4,
			output: 1,
		},
		{
			input:  6,
			output: 2,
		},
		{
			input:  5000,
			output: 5,
		},
		{
			input:  10000000000,
			output: 10,
		},
	}

	for _, td := range testData {
		res := CountPrime(td.input)
		if res != td.output {
			t.Errorf(fmt.Sprintf("For input %v - Expected : %v, Actual : %v", td.input, td.output, res))
		}
	}
}
