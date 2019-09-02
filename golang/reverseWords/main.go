package main

import "fmt"

func main() {
	fmt.Println(reverseWords("j k"))
}

//给定一个字符串，逐个翻转字符串中的每个单词。
//
//
//
//示例 1：
//
//输入: "the sky is blue"
//输出: "blue is sky the"

//无空格字符构成一个单词。
//输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
//如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

func reverseWords(s string) string {
	res := ""
	space := false
	start := len(s)
	end := len(s)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			space = true
			if end > start {
				res += s[start:end]
			}
			end = i
			start--
		} else {
			if space && len(res) > 0 {
				res += " "
			}
			space = false
			start--
		}
	}
	if end > start {
		res += s[start:end]
	}
	return res
}
