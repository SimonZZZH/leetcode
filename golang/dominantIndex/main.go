package main

import "fmt"

func main() {
	fmt.Println(dominantIndex([]int{0, 1}))
}

//在一个给定的数组nums中，总是存在一个最大元素 。
//
//查找数组中的最大元素是否至少是数组中每个其他数字的两倍。
//
//如果是，则返回最大元素的索引，否则返回-1。
func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var max, max2, index int
	if nums[0] > nums[1] {
		max = nums[0]
		max2 = nums[1]
		index = 0
	} else {
		max = nums[1]
		max2 = nums[0]
		index = 1
	}
	for i := 2; i < len(nums); i++ {
		if nums[i] > max {
			max2 = max
			max = nums[i]
			index = i
		} else if nums[i] > max2 {
			max2 = nums[i]
		}
	}
	if max >= max2*2 {
		return index
	}
	return -1
}
