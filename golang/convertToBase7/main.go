package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(convertToBase7(7))
}

//给定一个整数，将其转化为7进制，并以字符串形式输出。
func convertToBase7(num int) string {
	res := ""
	for math.Abs(float64(num)) >= 7 {
		k := num % 7
		res = strconv.Itoa(int(math.Abs(float64(k)))) + res
		num /= 7
	}
	res = strconv.Itoa(num) + res
	return res
}
