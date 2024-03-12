package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	result_len := len(nums) - k + 1
	result := make([]int, 0, result_len)

	for start := 0; start < result_len; start++ {
		slices := nums[start:k]
		k++
		max := slices[0]
		for _, v := range slices {
			if v > max {
				max = v
			}
		}
		result = append(result, max)
	}

	return result
}

func test1() {
	var nums = []int{1, 3, -1, -3, 5, 3, 6, 7}
	var k = 3
	fmt.Println(maxSlidingWindow(nums, k))
}

func test2() {
	var nums = []int{7, 2, 4}
	var k = 2
	fmt.Println(maxSlidingWindow(nums, k))
}

func main() {
	//test1()
	test2()
}
