package main

import "fmt"

func searchInsert(nums []int, target int) int {
	i := 0
	j := len(nums)
	for i < j {
		m := (i + j) / 2
		if nums[m] < target {
			i = m + 1
		} else if target < nums[m] {
			j = m
		} else {
			return m
		}
	}
	return i
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(nums, target))

	target2 := 2
	fmt.Println(searchInsert(nums, target2))

	target3 := 7
	fmt.Println(searchInsert(nums, target3))

	nums2 := []int{1, 3}
	target4 := 1
	fmt.Println(searchInsert(nums2, target4))
}
