package main

import (
	"fmt"
	"math/rand"
)

func main() {
	so := Constructor([]int{1, 2, 3, 4})
	fmt.Println(so.Reset())
	fmt.Println(so.Shuffle())
}

//打乱一个没有重复元素的数组。
// 以数字集合 1, 2 和 3 初始化数组。
//int[] nums = {1,2,3};
//Solution solution = new Solution(nums);
//
// 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。
//solution.shuffle();
//
// 重设数组到它的初始状态[1,2,3]。
//solution.reset();
//
// 随机返回数组[1,2,3]打乱后的结果。
//solution.shuffle();

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	var s Solution
	s.nums = nums
	return s
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.nums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	tmp := make([]int, len(this.nums))
	copy(tmp, this.nums)
	res := make([]int, len(this.nums))
	for x := 0; x < len(this.nums); x++ {
		index := rand.Intn(len(tmp))
		res[x] = tmp[index]
		tmp = append(tmp[:index], tmp[index+1:]...)
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
