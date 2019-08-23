//给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
//
//说明:
//
//    初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
//    你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
//
//示例:
//
//输入:
//nums1 = [1,2,3,0,0,0], m = 3
//nums2 = [2,5,6],       n = 3
//
//输出: [1,2,2,3,5,6]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/merge-sorted-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package main

import "fmt"

func main() {
	nums1 := []int{}
	nums2 := []int{}
	merge(nums1, 4, nums2, 4)
	fmt.Println(nums1)
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) == 0 || n == 0 {
		return
	}
	tmp := make([]int, 0)
	for x := 0; x < m; x++ {
		tmp = append(tmp, nums1[x])
	}
	//tmp := nums1[:m]
	//fmt.Println(tmp)
	indexTmp := 0
	index1 := 0
	index2 := 0
	for {
		//fmt.Println(tmp)
		if indexTmp < m && index2 < n && tmp[indexTmp] <= nums2[index2] {
			//fmt.Println(tmp[indexTmp], nums2[index2], "===", indexTmp)
			nums1[index1] = tmp[indexTmp]
			index1++
			indexTmp++
		} else if index2 < n {
			nums1[index1] = nums2[index2]
			index1++
			index2++
		} else if indexTmp < m {
			nums1[index1] = tmp[indexTmp]
			index1++
			indexTmp++
		}
		if index1 == len(nums1) {
			break
		}
	}

}
