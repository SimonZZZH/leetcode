package main

import "fmt"

func main() {
	mySlice := []int{1, 6, 5, 4, 3}
	fmt.Println(mySlice)
	rotate(mySlice, 2)
	fmt.Println(mySlice)
}

//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

//使用额外数组
func rotate(nums []int, k int) {
	if len(nums) <= 1 || k == 0 {
		return
	}
	numOfMove := k % len(nums)
	//fmt.Println(numOfMove, "-=-=-=")
	if numOfMove == 0 {
		return
	}
	tmp := make([]int, numOfMove)
	length := len(nums)
	copy(tmp, nums[length-numOfMove:])
	//fmt.Println(tmp, "====", nums)
	index := length - 1
	for x := length - numOfMove - 1; x >= 0; x-- {
		//fmt.Println(nums[x], nums)
		nums[index] = nums[x]

		index--
	}
	//copy(nums[:numOfMove], tmp)
	for x := 0; x < numOfMove; x++ {
		nums[x] = tmp[x]
	}
}

//不使用额外数组
func rotate2(nums []int, k int) {
	if len(nums) <= 1 || k == 0 {
		return
	}
	//移动次数，最多移动数组长度次
	numOfMove := k % len(nums)
	var tmp int
	for x := 0; x < numOfMove; x++ {
		tmp = nums[len(nums)-1]
		for index := len(nums) - 1; index > 0; index-- {
			nums[index] = nums[index-1]
		}
		nums[0] = tmp
	}
}
