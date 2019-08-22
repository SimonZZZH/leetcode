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
//请判断一个链表是否为回文链表。

type ListNode struct {
	Val  int
	Next *ListNode
}

//
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	tmp := head
	nums := make([]int, 0)
	//indexR := -1
	//indexL := -2
	for {
		nums = append(nums, tmp.Val)
		tmp = tmp.Next
		if tmp == nil {
			break
		}
	}
	for x := 0; x < len(nums)/2; x++ {
		if nums[x] != nums[len(nums)-1-x] {
			return false
		}
	}
	return true
}
