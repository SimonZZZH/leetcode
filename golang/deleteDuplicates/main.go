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

type ListNode struct {
	Val  int
	Next *ListNode
}

//给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
//删除排序链表中的val重复的节点

//优化,4ms
func deleteDuplicates2(head *ListNode) *ListNode {
	L := head
	R := head
	for R != nil && R.Next != nil {
		for R != nil && R.Next != nil && R.Val == R.Next.Val {
			R = R.Next
		}
		L.Next = R.Next
		R = R.Next
		L = R
	}
	return head
}

//12ms
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	L := head
	R := head.Next
	for {
		if R.Next == nil {
			if R.Val == L.Val {
				//fmt.Println(R.Val, "+++")
				L.Next = nil
				//fmt.Println(R.Next)
				return head
			} else {
				return head
			}
		} else {
			if R.Val == L.Val {
				R = R.Next
				L.Next = R
			} else {
				//L.Next = R
				R = R.Next
				L = L.Next
			}
		}
	}
}
