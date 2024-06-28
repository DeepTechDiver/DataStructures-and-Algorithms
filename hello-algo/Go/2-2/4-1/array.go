package main

import (
	"fmt"
	"math/rand"
)

func randomAccess(nums []int) (randomNum int) {
	// 在区间 [0, nums.length) 中随机抽取一个数字 作为返回值
	randomIndex := rand.Intn(len(nums))
	randomNum = nums[randomIndex]
	return
}

func insert(nums []int, num int, index int) {
	// 把索引 index 以及之后的所有元素向后移动一位
	for i := len(nums) - 1; i > index; i-- {
		nums[i] = nums[i-1]
	}
	// 将 num 赋给 index 处的元素
	nums[index] = num
}

func remove(nums []int, index int) {
	// 把索引 index 之后的所有元素向前移动一位
	for i := index; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
}

/* 遍历数组 */
func traverse(nums []int) {
	count := 0
	// 通过索引遍历数组
	for i := 0; i < len(nums); i++ {
		count += nums[i]
	}
	count = 0
	// 直接遍历数组元素
	for _, num := range nums {
		count += num
	}
	// 同时遍历数据索引和元素
	for i, num := range nums {
		count += nums[i]
		count += num
	}
}

/* 在数组中查找指定元素 */
func find(nums []int, target int) (index int) {
	index = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			index = i
			break
		}
	}
	return
}

/* 扩展数组长度 */
func extend(nums []int, enlarge int) []int {
	// 初始化一个扩展长度后的数组
	res := make([]int, len(nums)+enlarge)
	// 将原数组中的所有元素复制到新数组
	for i, num := range nums {
		res[i] = num
	}
	// 返回扩展后的新数组
	return res
}

func main() {
	var arr [5]int
	fmt.Println("arr =", arr)

	nums := []int{1, 3, 2, 5, 4}
	fmt.Println("nums =", nums)

	randomNum := randomAccess(nums)
	fmt.Println("在 nums 中获取随机元素", randomNum)

	insert(nums, 10, 3)

	traverse(nums)
	index := find(nums, 3)
	fmt.Println("在 nums 中查询元素 3 ，得到索引 =", index)

}
