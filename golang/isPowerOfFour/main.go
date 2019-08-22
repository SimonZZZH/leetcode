//给定一个整数 (32 位有符号整数)，请编写一个函数来判断它是否是 4 的幂次方。
package main

import "fmt"

func main() {
	fmt.Println(isPowerOfFour(9))
}

//不用特殊数字做掩码
func isPowerOfFour(num int) bool {
	if num < 1 { //1是4的0次方
		return false
	}
	if num&(num-1) == 0 { //2的幂
		//循环除4
		for num%4 == 0 {
			num /= 4
		}
		return num == 1
	}
	return false
}
