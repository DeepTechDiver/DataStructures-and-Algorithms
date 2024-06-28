package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(forLoopRecur(2))
}

/* 使用迭代模拟递归 */
func forLoopRecur(n int) int {
	// 使用一个显式的栈来模拟系统调用栈
	stack := list.New()
	res := 0
	// 递：递归调用
	for i := n; i > 0; i-- {
		// 通过“入栈操作”模拟“递”
		stack.PushBack(i)
	}
	// 归：返回结果
	for stack.Len() != 0 {
		// 通过“出栈操作”模拟“归”
		res += stack.Back().Value.(int)
		stack.Remove(stack.Back())
	}
	// res = 1+2+3+...+n
	return res
}
