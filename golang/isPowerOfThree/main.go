package main

import "fmt"

func main() {
	fmt.Println(isPowerOfThree(333))
}

//1162261467
//给定一个整数，写一个函数来判断它是否是 3 的幂次方。
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	if 1162261467%n == 0 {
		return true
	}
	return false
}

//循环
func isPowerOfThree2(n int) bool {
	if n <= 0 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}
