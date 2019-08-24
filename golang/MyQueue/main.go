package main

func main() {

}

//使用栈实现队列的下列操作：
//
//push(x) -- 将一个元素放入队列的尾部。
//pop() -- 从队列首部移除元素。
//peek() -- 返回队列首部的元素。
//empty() -- 返回队列是否为空。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/implement-queue-using-stacks
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//golang里面没有栈，随便写写吧
type MyQueue struct {
	data []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	var queue MyQueue
	queue.data = make([]int, 0)
	return queue
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.data = append(this.data, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	res := this.data[0]
	this.data = this.data[1:]
	return res
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	res := this.data[0]
	return res
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.data) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
