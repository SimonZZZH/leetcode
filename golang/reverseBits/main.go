package main

import "fmt"

func main() {
	var k uint32
	k = 236
	//fmt.Printf("%b\n", k)
	fmt.Printf("%b\n", reverseBits(k))
}

//颠倒给定的 32 位无符号整数的二进制位。
func reverseBits(num uint32) uint32 {
	var res uint32
	for x := 0; x < 32; x++ {
		if num%2 != 0 {
			res <<= 1
			res++
		} else {
			res <<= 1
		}
		num >>= 1
	}
	return res
}
