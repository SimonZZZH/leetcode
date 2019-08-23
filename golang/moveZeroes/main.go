package main

import "fmt"

func main() {
	tmp := []int{1, 2, 3, 0, 5, 6, 7}
	moveZeroes(tmp)
	fmt.Println(tmp)
}

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//双指针

func moveZeroes(nums []int) {
	var index int
	index = -1
	for x := 0; x < len(nums); x++ {
		if nums[x] == 0 {
			index = x
			break
		}
	}
	if index == -1 {
		return
	}
	//fmt.Println(index)
	for x := index + 1; x < len(nums); x++ {
		if nums[x] != 0 {
			nums[index], nums[x] = nums[x], nums[index]
			//fmt.Println(index)
			index++
		}
	}
}
