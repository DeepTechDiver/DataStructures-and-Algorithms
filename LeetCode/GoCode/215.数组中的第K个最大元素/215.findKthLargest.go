package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	// 快速选择算法
	left, right := 0, len(nums)-1
	for {
		pivotIndex := partition(nums, left, right)
		if pivotIndex == k-1 {
			return nums[pivotIndex]
		} else if pivotIndex < k-1 {
			left = pivotIndex + 1
		} else {
			right = pivotIndex - 1
		}
	}
}

// 普通快速排序
func QuickSort(array []int, left, right int) {
	if left < right {
		// 进行切分
		loc := partition(array, left, right)
		// 对左部分进行快排
		QuickSort(array, left, loc-1)
		// 对右部分进行快排
		QuickSort(array, loc+1, left)
	}
}

// 切分函数，返回切分元素的下标
func partition(arr []int, left, right int) int {
	//pivot := arr[right]
	//i := left
	//for j := left; j < right; j++ {
	//	if arr[j] >= pivot {
	//		arr[i], arr[j] = arr[j], arr[i]
	//		//swap(arr, i, j)
	//		i++
	//	}
	//}
	//arr[i], arr[right] = arr[right], arr[i]
	////swap(arr, i, right)
	//return i

	i := left + 1
	j := right
	for i < j {
		if arr[i] > arr[left] {
			arr[i], arr[j] = arr[j], arr[i]
			j--
		} else {
			i++
		}
	}
	if arr[i] >= arr[left] {
		i--
	}
	arr[left], arr[i] = arr[i], arr[left]
	return i
}

func swap(arr []int, i, j int) {
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4} // 示例数组
	k := 2                          // 要找的第 k 个最大元素
	fmt.Println("数组:", nums)
	fmt.Printf("第 %d 个最大元素是: %d\n", k, findKthLargest(nums, k))

	QuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
