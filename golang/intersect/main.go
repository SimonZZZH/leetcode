package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 3, 4}, []int{2, 2, 2, 3, 3, 3, 7, 8, 5}))
}

//给定两个数组，编写一个函数来计算它们的交集。
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	map1 := make(map[int]int)
	map2 := make(map[int]int)
	for _, v := range nums1 {
		if _, ok := map1[v]; ok {
			map1[v]++
		} else {
			map1[v] = 1
		}
	}
	for _, v := range nums2 {
		if _, ok := map2[v]; ok {
			map2[v]++
		} else {
			map2[v] = 1
		}
	}
	res := make([]int, 0)
	for k, v1 := range map1 {
		if v2, ok := map2[k]; ok {
			num := int(math.Min(float64(v1), float64(v2)))
			for x := 0; x < num; x++ {
				res = append(res, k)
			}
		}
	}
	return res
}
