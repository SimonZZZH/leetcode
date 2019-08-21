//给定一个二进制数组， 计算其中最大连续1的个数。
//
//示例 1:
//
//输入: [1,1,0,1,1,1]
//输出: 3
//解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.
//
//注意：
//
//    输入的数组只包含 0 和1。
//    输入数组的长度是正整数，且不超过 10,000。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/max-consecutive-ones
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 1, 1, 1}))
}

//func findMaxConsecutiveOnes(nums []int) int {
//	var i, j, res int
//	for i = 0; i < len(nums); i++ {
//		if nums[i] != 1 {
//			continue
//		}
//		tmp := 0
//		for j = i; j < len(nums); j++ {
//			if nums[j] == 1 {
//				tmp++
//			} else {
//				break
//			}
//		}
//		if tmp > res {
//			res = tmp
//		}
//	}
//	return res
//}

func findMaxConsecutiveOnes(nums []int) int {
	var res, tmp int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			tmp++
		} else {
			res = int(math.Max(float64(tmp), float64(res)))
			tmp = 0
		}
	}
	return int(math.Max(float64(tmp), float64(res)))
}
