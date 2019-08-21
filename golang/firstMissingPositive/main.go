package main

import "fmt"

func main() {
	fmt.Println(firstMissingPositive([]int{1, 1}))
}

//给定一个未排序的整数数组，找出其中没有出现的最小的正整数。
//你的算法的时间复杂度应为O(n)，并且只能使用常数级别的空间。
//桶排序
//未出现的最小正整数一定不会大于数组长度
func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	length := len(nums)
	for x := 0; x < length; x++ {
		for {
			if nums[x] > 0 && nums[x] <= length && nums[x] != x+1 && nums[nums[x]-1] != nums[x] {
				nums[nums[x]-1], nums[x] = nums[x], nums[nums[x]-1]
			} else {
				break
			}
		}
	}

	for x := 0; x < length; x++ {
		if nums[x] != x+1 {
			return x + 1
		}
	}
	return length + 1
}
