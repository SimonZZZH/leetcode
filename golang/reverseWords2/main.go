package main

import "fmt"

func main() {
	fmt.Println(reverseWords("asd ip jjk asd,k"))
}

//给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
//
//示例 1:
//
//输入: "Let's take LeetCode contest"
//输出: "s'teL ekat edoCteeL tsetnoc"
//
//注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。

func reverseWords(s string) string {
	var l, r int
	l = -1
	r = -1
	var res []byte
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			for k := r; k > l; k-- {
				res = append(res, s[k])
			}
			res = append(res, ' ')
			l = i
		}
		r++
	}
	for k := r; k > l; k-- {
		res = append(res, s[k])
	}
	return string(res)
}
