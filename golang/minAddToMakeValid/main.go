package main

import "fmt"

func main() {
	fmt.Println(minAddToMakeValid("())()((("))
}

//给定一个由 '(' 和 ')' 括号组成的字符串 S，
// 我们需要添加最少的括号（ '(' 或是 ')'，
// 可以在任何位置），以使得到的括号字符串有效。

//使用栈，将所有（）正确配对全部清除
func minAddToMakeValid(S string) int {
	if len(S) == 0 {
		return 0
	}
	stack := make([]byte, 0)
	for x := 0; x < len(S); x++ {
		if len(stack) > 0 {
			if stack[len(stack)-1] == '(' && S[x] == ')' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, S[x])
			}
		} else {
			stack = append(stack, S[x])
		}
	}
	return len(stack)
}
