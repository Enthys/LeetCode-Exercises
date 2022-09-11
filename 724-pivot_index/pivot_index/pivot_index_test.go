package pivot_index_test

import (
	"fmt"
	"pivot_index/pivot_index"
	"testing"
)

func TestPivotIndex(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{input: []int{1, 7, 3, 6, 5, 6}, expected: 3},
		{input: []int{1, 2, 3}, expected: -1},
		{input: []int{2, 1, -1}, expected: 0},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("en given '%+v' should receive '%d'", testCase.input, testCase.expected), func(t *testing.T) {
			res := pivot_index.PivotIndex(testCase.input)

			if res != testCase.expected {
				t.Errorf("did not get the expected result")
			}
		})
	}
}
