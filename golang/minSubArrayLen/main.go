//给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。
// 如果不存在符合条件的连续子数组，返回 0。
//
//示例:
//
//输入: s = 7, nums = [2,3,1,2,4,3]
//输出: 2
//解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
//
//进阶:
//
//如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-size-subarray-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package main

import "fmt"

func main() {
	fmt.Println(minSubArrayLen(11, []int{1, 2, 3, 4, 5}))
}

func minSubArrayLen(s int, nums []int) int {
	//区间结束index
	var end int
	//subArray长度
	var length int

	//确定一个初始条件
	for length = 1; length <= len(nums); length++ {
		if sum(nums[0:0+length]) >= s {
			end = length
			break
		}
	}
	if length > len(nums) {
		return 0
	}

	for length > 0 {
		for end <= len(nums) {
			if sum(nums[end-length+1:end]) >= s {
				length--
				break
			}
			end++
		}
		if end > len(nums) {
			break
		}
	}
	return length
}

func sum(nums []int) int {
	var res int
	for x := 0; x < len(nums); x++ {
		res += nums[x]
	}
	return res
}
