package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{}, 0, nil},
		{[]int{1}, 1, nil},
		{[]int{1, 2, 3, 4, 5}, 3, []int{0, 1}},
		{[]int{1, 2, 3, 4, 5}, 7, []int{2, 4}},
		{[]int{1, 2, 3, 4, 5}, 9, []int{3, 4}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{-1, -2, -3, -4, -5}, -3, []int{0, 1}},
		{[]int{-1, -2, -3, -4, -5}, -7, []int{2, 4}},
		{[]int{-1, -2, -3, -4, -5}, -9, []int{3, 4}},
		{[]int{0, 0, 0, 0, 0}, 0, []int{0, 1}},
	}

	for _, tc := range testCases {
		result := twoSum(tc.nums, tc.target)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("twoSum(%v, %v) = %v; expected %v", tc.nums, tc.target, result, tc.expected)
		}
	}
}
