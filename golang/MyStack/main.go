package main

func main() {

}

//使用队列实现栈的下列操作：
//
//push(x) -- 元素 x 入栈
//pop() -- 移除栈顶元素
//top() -- 获取栈顶元素
//empty() -- 返回栈是否为空
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/implement-stack-using-queues
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type MyStack struct {
	data []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	var stack MyStack
	stack.data = make([]int, 0)
	return stack
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.data = append(this.data, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	tmp := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	return tmp
	//if len(this.data) > 0 {
	//	tmp := this.data[len(this.data)-1]
	//	this.data = this.data[:len(this.data)-1]
	//	return tmp
	//} else {
	//	return 0
	//}
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.data[len(this.data)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.data) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
