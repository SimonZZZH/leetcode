//设计一个找到数据流中第K大元素的类（class）。注意是排序后的第K大元素，不是第K个不同的元素。
//
//你的 KthLargest 类需要一个同时接收整数 k 和整数数组nums 的构造器，它包含数据流中的初始元素。每次调用 KthLargest.add，返回当前数据流中第K大的元素。
//
//示例:
//
//int k = 3;
//int[] arr = [4,5,8,2];
//KthLargest kthLargest = new KthLargest(3, arr);
//kthLargest.add(3);   // returns 4
//kthLargest.add(5);   // returns 5
//kthLargest.add(10);  // returns 5
//kthLargest.add(9);   // returns 8
//kthLargest.add(4);   // returns 8
//
//说明:
//你可以假设 nums 的长度≥ k-1 且k ≥ 1。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/kth-largest-element-in-a-stream
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	kl := Constructor(2, []int{0})
	fmt.Println(kl.nums)
	fmt.Println(kl.Add(-1))
	fmt.Println(kl.nums)
}

type KthLargest struct {
	k    int
	nums []int
}

func Constructor(k int, nums []int) KthLargest {
	var kthLargest KthLargest
	kthLargest.k = k
	kthLargest.nums = make([]int, k)
	for i := 0; i < k; i++ {
		kthLargest.nums[i] = math.MinInt64
	}
	if len(nums) > 0 {
		sort.Ints(nums)
		length := len(kthLargest.nums)
		for i := len(nums) - 1; i >= 0; i-- {
			if length == 0 {
				break
			}
			kthLargest.nums[length-1] = nums[i]
			length--
		}
	}
	return kthLargest
}

func (this *KthLargest) Add(val int) int {

	if val > this.nums[0] {
		this.nums[0] = val
		//插入
		for i := 1; i < len(this.nums); i++ {
			if val > this.nums[i] {
				this.nums[i-1] = this.nums[i]
				this.nums[i] = val
			} else {
				this.nums[i-1] = val
				break
			}
		}
	}
	//
	return this.nums[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
