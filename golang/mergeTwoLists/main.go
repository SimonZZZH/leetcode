//将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//示例：
//
//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var tmp *ListNode
	if l1.Val < l2.Val {
		if l1.Next == nil {
			l1.Next = l2
		} else {
			tmp = l1.Next
			l1.Next = mergeTwoLists(tmp, l2)
		}
		return l1
	} else {
		if l2.Next == nil {
			l2.Next = l1
		} else {
			tmp = l2.Next
			l2.Next = mergeTwoLists(tmp, l1)
		}

	}
	return l2
}
