package main

import (
	"fmt"
	"testing"
)

func TestCountOverlap(t *testing.T) {
	type testType struct {
		bounds []int
		result int
	}
	tests := []testType{
		{[]int{5, 7, 7, 9}, 1},
		{[]int{2, 8, 3, 7}, 5},
		{[]int{6, 6, 4, 6}, 1},
		{[]int{2, 6, 4, 8}, 3},
		{[]int{10, 23, 21, 42}, 3},
		{[]int{10, 23, 11, 23}, 13},
		{[]int{10, 23, 10, 23}, 14},
	}
	for _, test := range tests {
		t.Run(fmt.Sprint(test.bounds), func(t *testing.T) {
			if count := countOverlap(test.bounds); count != test.result {
				t.Logf("%v, expected %d, got %d", test.bounds, test.result, count)
				t.Fail()
			}
			testReverse := testType{[]int{test.bounds[2], test.bounds[3], test.bounds[0], test.bounds[1]}, test.result}
			if count := countOverlap(testReverse.bounds); count != testReverse.result {
				t.Logf("Reverse: %v, expected %d, got %d", testReverse.bounds, testReverse.result, count)
				t.Fail()
			}
		})
	}

}
