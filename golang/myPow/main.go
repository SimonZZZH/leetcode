package main

import "fmt"

func main() {
	fmt.Println(myPow(2, 3))
	fmt.Println(myPow2(2, -2))
	fmt.Println(myPow2(99.9, 99))
}

//实现 pow(x, n) ，即计算 x 的 n 次幂函数。
//直接。。。。。超时
func myPow(x float64, n int) float64 {
	var res float64
	res = 1
	if n == 0 {
		return res
	}
	if x == 0 || x == 1 {
		return res
	}
	if n > 0 {
		for i := 0; i < n; i++ {
			res *= x
		}
	}
	if n < 0 {
		var xx float64
		if x < 0 {
			xx = -1 / float64(x)
		} else {
			xx = 1 / float64(x)
		}
		for i := 0; i < -n; i++ {
			res *= xx
		}
	}
	return res
}

//快速幂

func myPow2(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return fastPow(x, n)
}

func fastPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	half := fastPow(x, n/2)
	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}
