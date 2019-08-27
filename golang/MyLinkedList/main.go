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
	var list *MyLinkedList
	l := Constructor()
	list = &l
	//list := new(MyLinkedList)
	list.AddAtIndex(-1, 2)
	list.AddAtHead(1)
	list.AddAtTail(3)
	list.AddAtIndex(-1, 2)
	fmt.Println(list.Get(2))
	list.DeleteAtIndex(-1)
	fmt.Println(list.Get(1))
}

type MyLinkedList struct {
	length int
	head   *linked
}
type linked struct {
	val  int
	next *linked
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	var ll MyLinkedList
	ll.length = 0
	ll.head = nil
	return ll
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.length {
		return -1
	}
	tmp := this.head
	for i := 1; i <= index; i++ {
		tmp = tmp.next
	}
	return tmp.val
}

/** Add a node of value val before the first element of the linked list. After the insertion,
the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	lk := new(linked)
	lk.val = val
	tmp := this.head
	this.head = lk
	lk.next = tmp
	this.length++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	//tmp := this
	//for tmp.next != nil {
	//	tmp = tmp.next
	//}
	//var l MyLinkedList
	//l.val = val
	//tmp.next = &l
	tmp := this.head
	for tmp.next != nil {
		tmp = tmp.next
	}
	var ls linked
	ls.val = val
	ls.next = nil
	tmp.next = &ls
	this.length++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list,
the node will be appended to the end of linked list.
If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.length {
		return
	}
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	tmp := this.head
	for i := 1; i < index; i++ {
		tmp = tmp.next
	}
	next := tmp.next
	var ls linked
	ls.val = val
	ls.next = next
	tmp.next = &ls
	this.length++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	//fmt.Println("length", this.length)
	if index < 0 {
		return
	}
	if index > this.length-1 {
		return
	}
	if index == 0 {
		this.head = this.head.next
		this.length--
		return
	}
	tmp := this.head
	for i := 1; i < index; i++ {
		tmp = tmp.next
	}
	tmp.next = tmp.next.next
	this.length--
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

//func (this *MyLinkedList) Test() {
//	tmp := this
//	tmp.val = 88
//}
