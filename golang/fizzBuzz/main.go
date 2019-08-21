package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(20))
	fmt.Println(fizzBuzz2(20))
}

//写一个程序，输出从 1 到 n 数字的字符串表示。
//
//1. 如果 n 是3的倍数，输出“Fizz”；
//
//2. 如果 n 是5的倍数，输出“Buzz”；
//
//3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。

//172 ms	69.3 MB
func fizzBuzz(n int) []string {
	res := make([]string, n)
	for x := 1; x < n+1; x++ {
		tmp := x % 3
		tmp2 := x % 5
		if tmp == 0 && tmp2 == 0 {
			res[x-1] = "FizzBuzz"
		} else if tmp == 0 {
			res[x-1] = "Fizz"
		} else if tmp2 == 0 {
			res[x-1] = "Buzz"
		} else {
			res[x-1] = fmt.Sprintf("%d", x)
		}
	}
	return res
}

//改进
//140 ms	101.6 MB
func fizzBuzz2(n int) []string {
	res := make([]string, n)
	for x := 1; x < n+1; x++ {
		if x%15 == 0 {
			res[x-1] = "FizzBuzz"
		} else if x%3 == 0 {
			res[x-1] = "Fizz"
		} else if x%5 == 0 {
			res[x-1] = "Buzz"
		} else {
			res[x-1] = strconv.Itoa(x)
		}
	}
	return res
}

//224 ms	23 MB
func fizzBuzz3(n int) []string {
	res := make([]string, 0)
	for x := 1; x < n+1; x++ {
		if x%15 == 0 {
			res = append(res, "FizzBuzz")
		} else if x%3 == 0 {
			res = append(res, "Fizz")
		} else if x%5 == 0 {
			res = append(res, "Buzz")
		} else {
			res = append(res, strconv.Itoa(x))
		}
	}
	return res
}
