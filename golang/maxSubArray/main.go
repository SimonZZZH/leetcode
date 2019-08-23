//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
//示例:
//
//输入: [-2,1,-3,4,-1,2,1,-5,4],
//输出: 6
//解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
//
//进阶:
//
//如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-subarray
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -1, 6}))
}

//动态规划
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	for x := 1; x < len(nums); x++ {
		nums[x] += int(math.Max(float64(nums[x-1]), 0))
		maxSum = int(math.Max(float64(maxSum), float64(nums[x])))
	}
	return maxSum
}

//超时
func maxSubArray2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var maxSum int
	maxSum = nums[0]
	//var subArray []int
	for x := 0; x < len(nums); x++ {
		for y := x + 1; y <= len(nums); y++ {
			if sumOfArray(nums[x:y]) > maxSum {
				maxSum = sumOfArray(nums[x:y])
				//subArray = nums[x:y]
			}
		}
	}
	return maxSum
}

func sumOfArray(subArray []int) int {
	var sum int
	for x := 0; x < len(subArray); x++ {
		sum += subArray[x]
	}
	return sum
}
