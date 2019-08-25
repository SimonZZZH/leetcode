package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElement([]int{1, 2, 3}, []int{1, 2, 3, 4, 5, 6, 7}))
}

//
//给定两个没有重复元素的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。找到 nums1 中每个元素在 nums2 中的下一个比其大的值。
//
//nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出-1。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/next-greater-element-i
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//用map保存nums2的值对应的下表，
//遍历nums1。。。。。。
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	indexMap := make(map[int]int)
	for x := 0; x < len(nums2); x++ {
		indexMap[nums2[x]] = x
	}
	for x := 0; x < len(nums1); x++ {
		res[x] = -1
		for index := indexMap[nums1[x]] + 1; index < len(nums2); index++ {
			if nums2[index] > nums1[x] {
				res[x] = nums2[index]
				break
			}
		}
	}
	return res
}
