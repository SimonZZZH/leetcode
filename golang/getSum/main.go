//不使用运算符 + 和 - ​​​​​​​，计算两整数 ​​​​​​​a 、b ​​​​​​​之和。
//
//示例 1:
//
//输入: a = 1, b = 2
//输出: 3
//
//示例 2:
//
//输入: a = -2, b = 3
//输出: 1
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sum-of-two-integers
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

import "fmt"

func main() {
	fmt.Println(getSum(11, 0))
}

//异或为不进位加法
//&左移一位为进位
func getSum(a int, b int) int {
	var sum, carry int //进位
	sum = a ^ b
	carry = (a & b) << 1
	//fmt.Printf("%b,---%b\n", sum, carry)
	//直到进位为0
	for carry != 0 {
		tmp := sum
		sum ^= carry
		//fmt.Printf("%b***\n", sum)
		carry = (carry & tmp) << 1
		//fmt.Printf("%b===\n", carry)
	}
	return sum
}
