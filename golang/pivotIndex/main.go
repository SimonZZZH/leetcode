package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{-1, -1, -1, 0, 1, 1}))
	fmt.Println(pivotIndex([]int{-1, -1, -1, -1, -1, -1}))
}

//给定一个整数类型的数组 nums，请编写一个能够返回数组“中心索引”的方法。
//
//我们是这样定义数组中心索引的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。
//
//如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。

//
func pivotIndex(nums []int) int {
	if len(nums) < 1 {
		return -1
	}
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	lSum, rSum := 0, sum-nums[0]
	for i := 0; i < len(nums); i++ {
		if lSum == rSum {
			return i
		}
		if i == len(nums)-1 {
			break
		}
		lSum += nums[i]
		rSum -= nums[i+1]
	}
	return -1
}
