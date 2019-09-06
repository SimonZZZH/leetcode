//给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False。
//
//说明：不要使用任何内置的库函数，如  sqrt。
//
//示例 1：
//
//输入：16
//输出：True
//
//示例 2：
//
//输入：14
//输出：False
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-perfect-square
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 1000000; i++ {
		tmp := rand.Intn(9999)
		if tmp == 0 || tmp == 1 {
			continue
		}
		if !isPerfectSquare(tmp * tmp) {
			fmt.Println("error1", tmp)
			return
		}
		if isPerfectSquare(tmp*tmp + 1) {
			fmt.Println("error2", tmp)
			return
		}
	}
	fmt.Println("pass")
}
func isPerfectSquare(num int) bool {
	if num < 0 {
		return false
	}
	if num == 0 || num == 1 {
		return true
	}
	tmp := num % 10
	if tmp == 2 || tmp == 3 || tmp == 7 || tmp == 8 {
		return false
	}
	//二分法
	left := 0
	right := num
	res := num / 2
	for right > left+1 {
		res = (right-left)/2 + left
		//fmt.Println(res)
		if num == res*res {
			return true
		}
		if res*res > num {
			right = res
		} else {
			left = res
		}
	}
	return false
}
