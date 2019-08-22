package main

import "fmt"

//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
//说明：本题中，我们将空字符串定义为有效的回文串。
func main() {
	fmt.Println(isPalindrome(""))
}

func isPalindrome(s string) bool {
	var tmp []byte
	for x := 0; x < len(s); x++ {
		if s[x] >= 'A' && s[x] <= 'Z' {
			tmp = append(tmp, s[x]+0x20)
		} else if (s[x] >= 'a' && s[x] <= 'z') || (s[x] >= '0' && s[x] <= '9') {
			tmp = append(tmp, s[x])
		} else {
			continue
		}
	}
	lIndex := 0
	rIndex := len(tmp) - 1
	for lIndex < rIndex {
		if tmp[lIndex] != tmp[rIndex] {
			return false
		}
		rIndex--
		lIndex++
	}
	return true
}
