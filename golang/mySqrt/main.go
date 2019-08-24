//实现 int sqrt(int x) 函数。
//
//计算并返回 x 的平方根，其中 x 是非负整数。
//
//由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
//
//示例 1:
//
//输入: 4
//输出: 2
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sqrtx
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mySqrt(8))
}

//牛顿迭代法
func mySqrt(x int) int {
	if x <= 0 {
		return 0
	}
	var res float64
	res = float64(x)
	for {
		last := res
		res = (res + float64(x)/res) / 2
		if math.Abs(res-last) < 0.1 {
			return int(res)
		}
	}
}
