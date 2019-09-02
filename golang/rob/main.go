package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(rob([]int{2, 1, 1, 2}))
}

//你是一个专业的小偷，计划偷窃沿街的房屋。
// 每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
// 如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//
//给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/house-robber
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//动态规划
//当前房间能获取的最大值dp[n]=MAX(dp[n-1],dp[n-2]+num)
//类似爬楼梯题目
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return int(math.Max(float64(nums[0]), float64(nums[1])))
	}
	var max int
	tmp1 := nums[0]
	tmp2 := int(math.Max(float64(nums[0]), float64(nums[1])))
	for x := 2; x < len(nums); x++ {
		max = int(math.Max(float64(tmp2), float64(tmp1+nums[x])))
		tmp1 = tmp2
		tmp2 = max

	}
	return max
}
