//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//
//示例 1:
//
//输入: 123
//输出: 321
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(3336))
}

func reverse(x int) int {
	xx := int(math.Abs(float64(x)))
	num := make([]byte, 0)
	var out int
	for {
		if xx < 10 {
			num = append(num, byte(xx))
			break
		}
		num = append(num, byte(xx%10))
		xx = xx / 10
	}
	//fmt.Println(num)
	for x := 0; x < len(num); x++ {
		out = out + int(int(num[x])*int(math.Pow10(len(num)-1-x)))
	}
	//处理正负
	if x < 0 {
		out = out * -1
	}
	if out > math.MaxInt32 {
		return 0
	} else if out < math.MinInt32 {
		return 0
	} else {
		return out
	}
}
