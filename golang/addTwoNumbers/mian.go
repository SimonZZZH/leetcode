//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//示例：
//
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/add-two-numbers
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//该解法在超出int64范围是会溢出
//需要自己构建加法运算

package main

import "fmt"

func main() {

	//fmt.Println(9 % 10)
	l1 := getList(10000001)
	l2 := getList(564)
	tmp := l1
	for {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
		if tmp == nil {
			break
		}
	}
	fmt.Println("end!!")

	ln := addTwoNumbers(l1, l2)
	tmp2 := ln
	for {
		fmt.Println(tmp2.Val)
		tmp2 = tmp2.Next
		if tmp2 == nil {
			break
		}
	}
	fmt.Println("end!!!")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	LN := new(ListNode)
	temp := l1
	count := 1
	var val1 int
	var val2 int
	for {
		val1 = val1 + temp.Val*count
		count = count * 10
		temp = temp.Next
		if temp == nil {
			break
		}
	}
	temp = l2
	count = 1
	for {
		val2 = val2 + temp.Val*count
		count = count * 10
		temp = temp.Next
		if temp == nil {
			break
		}
	}
	val := val1 + val2
	//fmt.Println(val, val1, val2)
	LN = getList(val)
	return LN
}

func getList(val int) *ListNode {
	ln := new(ListNode)
	ln.Val = val % 10
	if val < 10 {
		ln.Next = nil
		return ln
	}
	num := val / 10
	ln.Next = getList(num)
	return ln
}
