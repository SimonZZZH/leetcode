package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("a"))
}

//给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
//
//在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。
//
//注意:
//假设字符串的长度不会超过 1010。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-palindrome
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func longestPalindrome(s string) int {
	tmp := make(map[byte]int)
	for x := 0; x < len(s); x++ {
		if tmp[s[x]] > 0 {
			tmp[s[x]]++
		} else {
			tmp[s[x]] = 1
		}
	}
	var count int
	var once bool
	var odd bool
	for _, v := range tmp {
		if v == 1 {
			//出现1次
			once = true
		} else if v%2 != 0 {
			//奇数次
			odd = true
			count = count + v - 1
		} else {
			//偶数次
			count += v
		}
	}
	if once || odd {
		count++
	}
	return count
}
