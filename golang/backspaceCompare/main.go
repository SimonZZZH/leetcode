package main

import "fmt"

func main() {
	fmt.Println(backspaceCompare("asd#####aaa#", "kk##aaa"))
}

//给定 S 和 T 两个字符串，当它们分别被输入到空白的文本编辑器后，判断二者是否相等，并返回结果。 # 代表退格字符。

//栈的操作方式
func backspaceCompare(S string, T string) bool {
	ss := getStr(S)
	tt := getStr(T)
	return ss == tt
}

func getStr(str string) string {
	ss := make([]byte, 0)
	for x := 0; x < len(str); x++ {
		switch str[x] {
		case '#':
			if len(ss) > 0 {
				ss = ss[:len(ss)-1]
			}
		default:
			ss = append(ss, str[x])
		}
	}
	return string(ss)
}
