package main

import "fmt"

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

//找到两个链表相交的起始点
//循环两次，消除长度差

//执行用时 :108 ms, 在所有 Go 提交中击败了12.11% 的用户
//内存消耗 :6.4 MB, 在所有 Go 提交中击败了95.98%的用户
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA := headA
	pB := headB
	count := 0
	for {
		if count == 2 {
			return nil
		}
		if pA == pB && pA != nil {
			return pA
		}
		pA = pA.Next
		pB = pB.Next
		if pA == nil {
			count++
			pA = headB
		}
		if pB == nil {
			pB = headA
		}
	}
}

//map
//执行用时 :68 ms, 在所有 Go 提交中击败了46.48% 的用户
//内存消耗 :7.5 MB, 在所有 Go 提交中击败了6.03%的用户
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	nodeMap := make(map[*ListNode]bool)
	tmp := headA
	for tmp != nil {
		nodeMap[tmp] = true
		tmp = tmp.Next
		fmt.Println("q")
	}
	tmpB := headB
	for tmpB != nil {
		if _, ok := nodeMap[tmpB]; ok {
			return tmpB
		}
	}
	return nil
}
