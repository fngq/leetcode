package main

// https://leetcode.com/problems/remove-duplicates-from-sorted-array
func RemoveDuplicates(nums []int) int {
	i := 0
	for j := range nums {
		if nums[j] != nums[i] {
			i += 1
			nums[i] = nums[j]
		}
	}
	i += 1
	return i
}

func main() {

}
