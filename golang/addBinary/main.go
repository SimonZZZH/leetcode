//给定两个二进制字符串，返回他们的和（用二进制表示）。
//
//输入为非空字符串且只包含数字 1 和 0。
//
//示例 1:
//
//输入: a = "11", b = "1"
//输出: "100"
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/add-binary
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

import (
	"fmt"
)

func main() {
	fmt.Println(addBinary("1111", "1111"))
}

func addBinary(a string, b string) string {
	var res []byte
	var tmp byte //进位
	numMap := map[byte]byte{'0': 0, '1': 1}
	if len(a) < len(b) {
		res = make([]byte, len(b)+1)
		aa := ""
		for x := 0; x < len(b)-len(a); x++ {
			aa += "0"
		}
		a = aa + a
	} else {
		res = make([]byte, len(a)+1)
		bb := ""
		for x := 0; x < len(a)-len(b); x++ {
			bb += "0"
		}
		b = bb + b
	}
	fmt.Println(a, b)
	for x := len(a) - 1; x >= 0; x-- {
		switch numMap[a[x]] + numMap[b[x]] + tmp {
		case 0:
			res[x+1] = '0'
			tmp = 0
		case 1:
			res[x+1] = '1'
			tmp = 0
		case 2:
			res[x+1] = '0'
			tmp = 1
		case 3:
			res[x+1] = '1'
			tmp = 1
		}
	}
	res[0] = tmp + '0'
	if res[0] != '0' {
		return string(res)
	} else {
		return string(res[1:])
	}
}

//
//func addBinary(a string, b string) string {
//	aa, _ := strconv.ParseUint(a, 2, 64)
//	bb, _ := strconv.ParseUint(b, 2, 64)
//	cc := aa + bb
//	return strconv.FormatUint(cc, 2)
//}
