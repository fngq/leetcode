package main

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	var i, j int
	// merge two array into one sorted array
	merge := make([]int, len(nums1)+len(nums2))
	for m := range merge {
		if i >= len(nums1) {
			merge[m] = nums2[j]
			j++
			continue
		}
		if j >= len(nums2) {
			merge[m] = nums1[i]
			i++
			continue
		}
		if nums1[i] < nums2[j] {
			merge[m] = nums1[i]
			i++
			continue
		} else {
			merge[m] = nums2[j]
			j++
			continue
		}
	}
	if (len(nums1)+len(nums2))%2 == 0 {
		return float64(merge[len(merge)/2-1]+merge[len(merge)/2]) / 2
	}
	return float64(merge[len(merge)/2])
}
