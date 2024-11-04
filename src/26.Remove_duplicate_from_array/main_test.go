package main

import (
	"log"
	"testing"
)

func Test_remove_duplicate(t *testing.T) {
	testCase := []struct {
		nums   []int
		expect []int
	}{
		{[]int{1, 1, 2}, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
	}
	for _, tc := range testCase {
		log.Println("test case", tc.nums)
		ret := RemoveDuplicates(tc.nums)
		if ret != len(tc.expect) {
			t.Errorf("wrong length %v,expect %v", ret, len(tc.expect))
		}
		for i := 0; i < ret; i++ {
			if tc.nums[i] != tc.expect[i] {
				t.Errorf("wrong at %v", i)
			}
		}
	}

}
