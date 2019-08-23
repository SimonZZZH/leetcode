package main

import "fmt"

func main() {
	fmt.Println(isValid("(([]){})"))
}

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//
//注意空字符串可被认为是有效字符串。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-parentheses
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//栈的方法
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 > 0 {
		return false
	}

	myMap := map[byte]byte{'(': ')', '[': ']', '{': '}'}
	stack := make([]byte, 0)
	for x := 0; x < len(s); x++ {
		if _, ok := myMap[s[x]]; ok {
			stack = append(stack, s[x])
		} else {
			if len(stack) == 0 {
				return false
			}
			if myMap[stack[len(stack)-1]] == s[x] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
