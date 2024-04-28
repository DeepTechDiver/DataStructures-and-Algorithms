package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}
	// 偶数和奇数 当为奇数时中位数不影响判断，所以可以通过 除10 去除反转后的中位数，剩余结果与进行对比
	return x == revertedNumber || x == revertedNumber/10
}

func main() {
	testNums := []int{121, -121, 10}
	for _, num := range testNums {
		fmt.Println(isPalindrome(num))
	}
}
