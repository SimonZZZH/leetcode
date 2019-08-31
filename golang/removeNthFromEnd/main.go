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

//给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	tmp := make([]*ListNode, 0)
	p := head
	p2 := head.Next
	for {
		tmp = append(tmp, p)
		p = p2
		if p == nil {
			break
		}
		p2 = p.Next
	}
	if n == len(tmp) {
		return head.Next
	} else if n > 1 {
		tmp[len(tmp)-n-1].Next = tmp[len(tmp)-n+1]
	} else {
		tmp[len(tmp)-n-1].Next = nil
	}
	return head
}
