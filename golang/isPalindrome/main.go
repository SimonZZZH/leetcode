//回文数字判断
package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(11111111))

}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	xSlice := make([]byte, 0)
	//var tmp int
	for {
		xSlice = append(xSlice, byte(x%10))
		x = x / 10
		if x == 0 {
			break
		}
	}
	for x := 0; x < len(xSlice); x++ {
		if xSlice[x] != xSlice[len(xSlice)-1-x] {
			return false
		} //25652
		if x >= len(xSlice)-1-x {
			break
		}
	}
	return true
}
