//给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，
// 使得 nums [i] = nums [j]，并且 i 和 j 的差的绝对值最大为 k。
//此处描述有误，应该是不大于k
//示例 1:
//
//输入: nums = [1,2,3,1], k = 3
//输出: true
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/contains-duplicate-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

import "fmt"

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
}
func containsNearbyDuplicate(nums []int, k int) bool {
	numsMap := make(map[int]int) //存放索引在nums中最新出现的位置
	//resMap := make(map[int]int)  //存放索引数字下标的最大差值
	for i := 0; i < len(nums); i++ {
		if val, ok := numsMap[nums[i]]; ok {
			//再次出现
			if i-val <= k {
				return true
			}
		}
		//更新位置
		numsMap[nums[i]] = i
	}
	//for _, val := range resMap {
	//	if val == k {
	//		return true
	//	}
	//}
	return false
}
