package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(28))
}

//给定一个正整数，返回它在 Excel 表中相对应的列名称。
//
//例如，
//
//1 -> A
//2 -> B
//3 -> C
//...
//26 -> Z
//27 -> AA
//28 -> AB
//...
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/excel-sheet-column-title
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//因为没有0，所以。。。。还真不太好解释
//每次n-1取余，余数加1就是当前位，可避免出现余数为0，

func convertToTitle(n int) string {
	var res string
	resSlice := make([]byte, 0)
	var quotient int   //商
	var remainder byte //余数
	for {
		if n == 0 {
			break
		}
		quotient = (n - 1) / 26
		remainder = byte((n-1)%26 + 1)
		resSlice = append(resSlice, remainder+0x40)
		n = quotient
	}
	for x := len(resSlice) - 1; x >= 0; x-- {
		res += string(resSlice[x])
	}
	return res
}
