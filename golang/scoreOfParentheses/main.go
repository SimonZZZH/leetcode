package main

import "fmt"

func main() {
	fmt.Println(scoreOfParentheses(""))
}

//给定一个平衡括号字符串 S，按下述规则计算该字符串的分数：
//
//() 得 1 分。
//AB 得 A + B 分，其中 A 和 B 是平衡括号字符串。
//(A) 得 2 * A 分，其中 A 是平衡括号字符串。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/score-of-parentheses
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//乘法分配律
//栈
func scoreOfParentheses(S string) int {
	stack := uint(0) //记录乘法深度
	var res int
	for x := 0; x < len(S); x++ {
		if S[x] == '(' {
			stack++ //匹配到左括号，入栈
		} else if S[x] == ')' && S[x-1] == '(' {
			//匹配到（）
			stack--
			res = res + (1 << stack)
		} else {
			//匹配到右括号，出栈
			stack--
		}
	}
	return res
}
