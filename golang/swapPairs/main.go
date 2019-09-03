package main

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := head.Next
	node1 := head
	node2 := head.Next
	for {
		if node2.Next == nil {
			node2.Next = node1
			node1.Next = nil
			break
		} else if node2.Next.Next == nil {
			tmp := node2.Next
			node2.Next = node1
			node1.Next = tmp
			break
		} else {
			tmp := node2.Next
			node2.Next = node1
			node1.Next = tmp.Next
			node1 = tmp
			node2 = tmp.Next
		}
	}
	return res
}
