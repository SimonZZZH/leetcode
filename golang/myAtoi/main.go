//请你来实现一个 atoi 函数，使其能将字符串转换成整数。
//
//首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。
//
//当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，作为该整数的正负号；
// 假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。
//
//该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。
//
//注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。
//
//在任何情况下，若函数不能进行有效的转换时，请返回 0。
//
//说明：
//
//假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231,  231 − 1]。
// 如果数值超过这个范围，qing返回  INT_MAX (231 − 1) 或 INT_MIN (−231) 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/string-to-integer-atoi
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("     -0000009"))

}

func myAtoi(str string) int {
	numberSlice := make([]byte, 0)
	//去空格
	for x := 0; x < len(str); x++ {
		if str[x] != ' ' {
			str = str[x:]
			break
		}
		if x == len(str)-1 {
			str = ""
			break
		}
	}
	if len(str) == 0 {
		return 0
	}
	//确定正负
	switch str[0] {
	case '-', '+':
		numberSlice = append(numberSlice, str[0])
		str = str[1:]
	default:
		numberSlice = append(numberSlice, '+')
	}
	//去0
	for x := 0; x < len(str); x++ {
		if str[x] != '0' {
			str = str[x:]
			break
		}
		if x == len(str)-1 {
			str = ""
			break
		}
	}
	if len(str) == 0 {
		return 0
	}
	//存到数组
	for x := 0; x < len(str); x++ {
		if str[x] >= 0x30 && str[x] <= 0x39 {
			numberSlice = append(numberSlice, str[x]-0x30)
		} else {
			break
		}
	}
	if len(numberSlice) > 11 {
		switch numberSlice[0] {
		case '-':
			return math.MinInt32
		default:
			return math.MaxInt32
		}
	}
	var number int64
	for x := 1; x < len(numberSlice); x++ {
		number = number + int64(int64(numberSlice[x])*int64(math.Pow10((len(numberSlice)-x-1))))
	}
	if numberSlice[0] == '-' {
		number = number * -1
	}
	if number > math.MaxInt32 {
		return math.MaxInt32
	} else if number < math.MinInt32 {
		return math.MinInt32
	} else {
		return int(number)
	}
}
