package main

import "fmt"

func main() {
	tmp := []int{1, 1, 2}
	fmt.Println(removeDuplicates(tmp), tmp, "=====")
}

//给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
//
//不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//优化
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	x := 0
	y := x + 1
	for y < len(nums) {
		if nums[x] != nums[y] {
			x++
			nums[x] = nums[y]
		}
		y++
	}
	return x + 1
}

//自己写的
//func removeDuplicates(nums []int) int {
//	if len(nums) < 2 {
//		fmt.Println("sss")
//		return len(nums)
//	}
//	x := 0
//	y := x + 1
//	count := 0
//	for x < len(nums) {
//		for y < len(nums)-count {
//			if nums[y] != nums[x] {
//				break
//			}
//			y++
//		}
//		//fmt.Println(x, y)
//		if y == len(nums) {
//			return x + 1
//		}
//		//数据移动
//		if y-x > 1 {
//			for k := 0; k < len(nums)-y-count; k++ {
//				nums[x+1+k] = nums[y+k]
//			}
//			count = count + (y - x - 1)
//		}
//		if x+count >= len(nums)-1 {
//			break
//		}
//		x++
//		y = x + 1
//	}
//	return x + 1
//}
