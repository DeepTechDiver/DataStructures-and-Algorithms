package main

import "fmt"

func combinationSum4(nums []int, target int) int {
	result := make([]int, target+1)
	result[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				result[i] += result[i-num]
			}
		}
	}
	return result[target]
}

func main() {
	nums := []int{1, 2, 3}
	target := 4
	fmt.Println(combinationSum4(nums, target))
}
