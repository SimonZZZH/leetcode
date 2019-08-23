package main

import "fmt"

func main() {
	sm := Constructor()
	p := &sm
	for x := 0; x < 10; x++ {
		p.Push(x)
	}
	fmt.Println(p.GetMin())
}

//设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
//
//push(x) -- 将元素 x 推入栈中。
//pop() -- 删除栈顶的元素。
//top() -- 获取栈顶元素。
//getMin() -- 检索栈中的最小元素。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/min-stack
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type MinStack struct {
	data   []int //栈
	minNum []int //辅助栈，用于存放最小值
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var ms MinStack
	ms.data = make([]int, 0)
	ms.minNum = make([]int, 0)
	return ms
}

func (this *MinStack) Push(x int) {
	if len(this.data) != 0 {
		if x <= this.minNum[len(this.minNum)-1] {
			this.minNum = append(this.minNum, x)
		}
	} else {
		this.minNum = append(this.minNum, x)
	}
	this.data = append(this.data, x)
}

func (this *MinStack) Pop() {
	if this.data[len(this.data)-1] == this.minNum[len(this.minNum)-1] {
		this.minNum = this.minNum[:len(this.minNum)-1]
	}
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.minNum[len(this.minNum)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
