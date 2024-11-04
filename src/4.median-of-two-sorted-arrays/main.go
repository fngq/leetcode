package main

import (
	"log"
)

/*
4. 寻找两个正序数组的中位数
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
请你找出并返回这两个正序数组的 中位数 。
leetcode :https://leetcode.com/problems/median-of-two-sorted-arrays/description/
*/

/*
This is a general solution for problems like "find thd Kth number of an arry"
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)
	left := (m + n + 1) / 2
	right := (n + m + 2) / 2

	return float64(findKthMin(nums1, 0, nums2, 0, left)+findKthMin(nums1, 0, nums2, 0, right)) / 2.0

}

func findKthMin(nums1 []int, start1 int, nums2 []int, start2 int, k int) int {
	l1 := len(nums1) - start1
	l2 := len(nums2) - start2
	if l1 > l2 {
		return findKthMin(nums2, start2, nums1, start1, k)
	}
	log.Println(l1, start1, ",", l2, start2, k)
	if l1 == 0 {
		return nums2[start2+k-1]
	}
	if k == 1 {
		return min(nums1[start1], nums2[start2])
	}
	//index访问，需要-1
	i := start1 + min(k/2, l1) - 1
	j := start2 + min(k/2, l2) - 1

	if nums1[i] < nums2[j] {
		return findKthMin(nums1, i+1, nums2, start2, k-(i-start1+1))
	} else {
		return findKthMin(nums1, start1, nums2, j+1, k-(j-start2+1))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
