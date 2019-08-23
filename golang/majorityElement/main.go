package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{1, 2, 3, 7, 7, 7, 7}))
}

//给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
//
//你可以假设数组是非空的，并且给定的数组总是存在众数。

func majorityElement(nums []int) int {
	numMap := make(map[int]int)
	for x := 0; x < len(nums); x++ {
		if numMap[nums[x]] > len(nums)/2-1 {
			return nums[x]
		} else if numMap[nums[x]] > 0 {
			numMap[nums[x]]++
		} else {
			numMap[nums[x]] = 1
		}
	}
	return 0
}
