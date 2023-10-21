package main

import (
	"testing"
)

type normalizeTestCase struct{
	input string
	desired string
}

func TestNormalize(t *testing.T){
	testcases := []struct{
		input string
		desired string
	}{
		{"8182009821", "8182009821"},
		{"(81)82009821", "8182009821"}, 
		{"81 8200 9821", "8182009821"},
	}
for _, tc := range testcases{
	t.Run(tc.input, func(t *testing.T){
		actual := normalize(tc.input)
	if actual != tc.desired{
		t.Errorf("We got %s; but want %s", actual, tc.desired)
			}
		})
	}
}