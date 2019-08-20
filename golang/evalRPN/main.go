//根据逆波兰表示法，求表达式的值。
//
//有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
//
//说明：
//
//整数除法只保留整数部分。
//给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
//示例 3：
//
//输入: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
//输出: 22
//解释:
//((10 * (6 / ((9 + 3) * -11))) + 17) + 5
//= ((10 * (6 / (12 * -11))) + 17) + 5
//= ((10 * (6 / -132)) + 17) + 5
//= ((10 * 0) + 17) + 5
//= (0 + 17) + 5
//= 17 + 5
//= 22
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}

func evalRPN(tokens []string) int {
	stack := make([]int, len(tokens))
	//指针
	index := -1
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+":
			tmp := stack[index-1] + stack[index]
			index--
			stack[index] = tmp
		case "-":
			tmp := stack[index-1] - stack[index]
			index--
			stack[index] = tmp
		case "*":
			tmp := stack[index-1] * stack[index]
			index--
			stack[index] = tmp
		case "/":
			tmp := stack[index-1] / stack[index]
			index--
			stack[index] = tmp
		default:
			index++
			stack[index], _ = strconv.Atoi(tokens[i])
		}
	}
	return stack[0]
}
