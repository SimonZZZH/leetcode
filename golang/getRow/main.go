//给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
//
//在杨辉三角中，每个数是它左上方和右上方的数的和。
//
//示例:
//
//输入: 3
//输出: [1,3,3,1]
//
//进阶：
//
//你可以优化你的算法到 O(k) 空间复杂度吗？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/pascals-triangle-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

import "fmt"

func main() {
	fmt.Println(getRow(4))
}

//返回杨辉三角的第rowIndex行
func getRow(rowIndex int) []int {

	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}
	res := make([]int, rowIndex+1)
	for x := 0; x <= rowIndex; x++ {
		res[x] = 1
	}
	tmp := 1
	for x := 2; x <= rowIndex; x++ {
		for y := 1; y < x; y++ {
			res[y], tmp = res[y]+tmp, res[y]
		}
	}
	return res
}
