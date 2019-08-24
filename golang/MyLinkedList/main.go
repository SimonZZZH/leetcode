//设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，
// next 是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。
// 假设链表中的所有节点都是 0-index 的。
//
//在链表类中实现这些功能：
//
//    get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
//    addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
//    addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
//    addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。
// 如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。
// 如果index小于0，则在头部插入节点。
//    deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/design-linked-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package main

import "fmt"

func main() {
	list := new(MyLinkedList)
	list.AddAtHead(3)
	fmt.Println(list)
}

type MyLinkedList struct {
	val  int
	next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	var ll MyLinkedList
	return ll
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	tmp := this
	for i := 1; i <= index; i++ {
		if tmp.next != nil {
			tmp = tmp.next
		} else {
			return -1
		}
	}
	return tmp.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	fmt.Println(this)
	tmp := this
	this = new(MyLinkedList)
	this.val = val
	this.next = tmp
	fmt.Println(this)
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	tmp := this
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = new(MyLinkedList)
	tmp.next.val = val
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	tmp := this
	if index <= 0 {
		this.AddAtHead(val)
	} else {
		for i := 1; i < index; i++ {
			if tmp.next == nil {
				return
			}
			tmp = tmp.next
		}
		tmp2 := tmp.next
		tmp.next = new(MyLinkedList)
		tmp.next.val = val
		tmp.next.next = tmp2
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	tmp := this
	if index == 0 {
		this = this.next
		return
	}
	for i := 1; i < index; i++ {
		if tmp.next == nil {
			return
		}
		tmp = tmp.next
	}
	if tmp.next != nil {
		tmp.next = tmp.next.next
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
