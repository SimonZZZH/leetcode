package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate2([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

//给定一个整数数组，判断是否存在重复元素。
//
//如果任何值在数组中出现至少两次，函数返回 true。如果数组中每个元素都不相同，则返回 false。

//
func containsDuplicate2(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	dataMap := make(map[int]byte)
	for x := 0; x < len(nums); x++ {
		if dataMap[nums[x]] == 1 {
			return true
		}
		dataMap[nums[x]] = 1
	}
	return false
}

//双指针
func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	//index := 0

	for index := 0; index < len(nums); index++ {
		for x := index + 1; x < len(nums); x++ {
			if nums[index] == nums[x] {
				return true
			}
		}
	}
	return false
}
