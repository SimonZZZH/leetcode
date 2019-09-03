package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strStr("aassddfg", "ss"))

}

//TODO 实现KMP算法
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
