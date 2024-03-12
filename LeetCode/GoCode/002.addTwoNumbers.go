package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node1 := ListNode{Val: 1, Next: nil}

	for k, v := range  {
		node1 := ListNode{Val: k, Next: l1}
		node2 := ListNode{Val: k, Next: l2}
	}
}



func main() {
	var l1 = []int{2, 4, 3}
	var l2 = []int{5, 6, 4}
	addTwoNumbers(l1, l2)

}
