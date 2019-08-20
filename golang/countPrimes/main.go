package main

import "fmt"

func main() {
	fmt.Println(countPrimes(20))
}

//统计所有小于非负整数 n 的质数的数量。

//线性筛
func countPrimes(n int) int {
	numFlag := make([]bool, n)
	for x := 2; x < n; x++ {
		numFlag[x] = true
	}
	//fmt.Println(numFlag)
	//return 0
	var count int
	for x := 2; x < n; x++ {
		if numFlag[x] {
			count++
			for k := 2; k*x < n; k++ {
				numFlag[x*k] = false
			}
		}
	}
	//fmt.Println(numFlag)
	return count
}
