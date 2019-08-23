package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("a"))
}

//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
func longestPalindrome(s string) string {
	var length int
	var str string
	for x := 0; x < len(s); x++ {
		for y := x + 1; y <= len(s); y++ {
			l := 0
			r := len(s[x:y]) - 1
			for l < r {
				if s[x:y][l] != s[x:y][r] {
					break
				}
				l++
				r--
			}
			if l >= r {
				if len(s[x:y]) > length {
					length = len(s[x:y])
					str = string(s[x:y])
				}
			}
		}
	}
	return str
}
