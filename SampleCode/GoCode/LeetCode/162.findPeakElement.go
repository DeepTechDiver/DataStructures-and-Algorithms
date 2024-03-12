package main

import "fmt"

func findPeakElement(nums []int) int {
	i, j := 0, len(nums)-1
	for i < j {
		m := (i + j) >> 1
		if nums[m+1] < nums[m] {
			j = m
		} else {
			i = m + 1
		}
	}
	return i
}

func main() {
	nums := []int{1}
	fmt.Println(findPeakElement(nums))
}
