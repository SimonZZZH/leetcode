//给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
//
//最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
//
//你可以假设除了整数 0 之外，这个整数不会以零开头。
//
//示例 1:
//
//输入: [1,2,3]
//输出: [1,2,4]
//解释: 输入数组表示数字 123。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/plus-one
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{0}))
}

func plusOne(digits []int) []int {
	tmp := 1
	for x := len(digits) - 1; x >= 0; x-- {
		if tmp+digits[x] >= 10 {
			digits[x] = (tmp + digits[x]) % 10
			//tmp = 1
		} else {
			digits[x] += tmp
			tmp = 0
			break
		}
	}
	switch tmp {
	case 0:
		return digits
	case 1:
		res := append([]int{1}, digits...)
		return res
	}
	return nil
}
