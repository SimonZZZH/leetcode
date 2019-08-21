package main

import "fmt"

func main() {
	var num uint32
	num = 66
	fmt.Printf("%b\n", num)
	fmt.Println(hammingWeight(num))
}

//编写一个函数，输入是一个无符号整数，返回其二进制表达式中数字位数为 ‘1’ 的个数（也被称为汉明重量）。

func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		if num%2 != 0 {
			count++
		}
		num = num >> 1
	}
	return count
}
