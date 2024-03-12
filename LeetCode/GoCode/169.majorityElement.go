package main

import "fmt"

// 带兵打仗法
func majorityElement1(nums []int) int {
	n := len(nums)
	count := 1
	most_num := nums[0] // i=0
	for i := 1; i < n; i++ {
		if nums[i] == most_num {
			count++
		} else {
			count--
			if count == 0 {
				most_num = nums[i]
				count = 1
			}
		}
	}
	return most_num
}

//摩尔投票法
func majorityElement2(nums []int) int {
	major := 0
	count := 0
	for _, num := range nums {
		if count == 0 {
			major = num
		}
		if major == num {
			count += 1
		} else {
			count -= 1
		}
	}
	return major
}

func main() {
	var nums = []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement1(nums))

	fmt.Println(majorityElement2(nums))
}
