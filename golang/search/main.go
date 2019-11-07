//给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
//
//
//示例 1:
//
//输入: nums = [-1,0,3,5,9,12], target = 9
//输出: 4
//解释: 9 出现在 nums 中并且下标为 4
//
//示例 2:
//
//输入: nums = [-1,0,3,5,9,12], target = 2
//输出: -1
//解释: 2 不存在 nums 中因此返回 -1
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-search
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package main

import (
	"fmt"
)

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}

func search(nums []int, target int) int {

	left := 0
	right := len(nums)
	index := right / 2
	if right == 0 {
		return -1
	}
	if nums[left] == target {
		return left
	}
	if nums[right-1] == target {
		return right - 1
	}
	for index > 0 {
		if nums[left+index] == target {
			return left + index
		}
		if nums[left+index] > target {
			right = left + index
			index = (right - left) / 2
			//fmt.Println(index, "ddd")
			continue
		}
		if nums[left+index] < target {
			left = left + index
			//if nums[left] == target {
			//	return left
			//}
			index = (right - left) / 2
			//fmt.Println(right, left, index, "sss")
			//time.Sleep(time.Second * 3)
			continue
		}
	}
	return -1
}
