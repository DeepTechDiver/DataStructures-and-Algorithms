package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 创建链表节点
	node1 := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 2, Next: nil}
	node3 := ListNode{Val: 3, Next: nil}
	// 连接链表节点
	node1.Next = &node2
	node2.Next = &node3

	// 打印链表
	printList(&node1) //1 -> 2 -> 3 -> nil
	// 1. 插入节点
	insertNode(&node3, 4)
	printList(&node1) //1 -> 2 -> 3 -> 4 -> nil
	// 2. 删除节点
	// reverseList(&node2)
	printList(&node3)
	// 3. 查找节点
	fmt.Println(findNode(&node1, 4))
	// 4. 反转链表
	// 5. 检测环

}

// 1. 插入节点
func insertNode(prev *ListNode, val int) {
	newNode := ListNode{Val: val, Next: prev.Next}
	prev.Next = &newNode
}

// 2. 删除节点
func deleteNode(prev *ListNode) {
	if prev.Next != nil {
		prev.Next = prev.Next.Next
	}
}

// 3. 查找节点
func findNode(head *ListNode, val int) *ListNode {
	current := head
	for current != nil {
		if current.Val == val {
			return current
		}
		current = current.Next
	}
	return nil
}

// 4. 反转链表
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

// 5. 检测环
func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true // 有环
		}
	}

	return false // 无环
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}
