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

package main

import (
	"fmt"
)

func main() {

	l1 := getList([]byte{1, 1})
	l2 := getList([]byte{9, 8})
	ln := addTwoNumbers(l1, l2)
	tmp := getByteSlice(ln)
	fmt.Println(tmp)
	fmt.Println("end!!!")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	LN := new(ListNode)
	num1 := getByteSlice(l1)
	num2 := getByteSlice(l2)
	var length int
	if len(num1) > len(num2) {
		tmp := make([]byte, len(num1)-len(num2))
		num2 = append(num2, tmp...)
	} else if len(num2) > len(num1) {
		tmp := make([]byte, len(num2)-len(num1))
		num1 = append(num1, tmp...)
	}
	length = len(num1)
	number := make([]byte, length+1)
	for x := 0; x < length; x++ {
		sum := num1[x] + num2[x]
		number[x+1] = (number[x] + sum) / 10
		number[x] = (number[x] + sum) % 10

	}
	//若最高位进位预留没用到，则删除
	if number[len(number)-1] == 0 {
		number = number[:len(number)-1]
	}
	LN = getList(number)
	return LN
}

//1536表示为[6,3,5,1]
func getList(val []byte) *ListNode {
	ln := new(ListNode)
	//最好加上判断字符串是否符合整数规则
	ln.Val = int(val[0])
	if len(val) == 1 {
		ln.Next = nil
		return ln
	}
	num := val[1:len(val)]
	ln.Next = getList(num)
	return ln
}

func getByteSlice(ln *ListNode) []byte {
	numSlice := make([]byte, 0)
	lnTmp := ln
	for {
		numSlice = append(numSlice, byte(lnTmp.Val))
		lnTmp = lnTmp.Next
		if lnTmp == nil {
			break
		}
	}
	return numSlice
}
