package main

import (
	"fmt"
	"strconv"
)

func countBits(n int) []int {
	result := []int{}
	for i := 0; i <= n; i++ {
		count := 0
		binary := strconv.FormatInt(int64(i), 2)
		for _, v := range binary {
			if string(v) == "1" {
				count++
			}
		}
		result = append(result, count)
	}
	return result
}

func countBits2(n int) []int {
	bits := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bits[i] = bits[i-highBit] + 1
	}
	return bits
}

func main() {
	n := 2
	fmt.Println(countBits(n))

	n2 := 5
	fmt.Println(countBits(n2))
	fmt.Println(countBits2(n2))
}
