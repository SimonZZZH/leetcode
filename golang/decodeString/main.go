//给定一个经过编码的字符串，返回它解码后的字符串。
//
//编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
//
//你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
//
//此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
//
//示例:
//
//s = "3[a]2[bc]", 返回 "aaabcbc".
//s = "3[a2[c]]", 返回 "accaccacc".
//s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".
//

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(decodeString("2[abc]3[cd]ef"))
}

func decodeString(s string) string {
	res := make([]byte, 0)
	stack := make([]byte, 0)    //反过来的结果
	tmpStack := make([]byte, 0) //保存当前一对括号内的字符串

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '[' {
			//匹配到完整括号，开始出栈，直到遇到']'
			tmp := make([]byte, 0)
			for i := len(stack) - 1; i >= 0; i-- {
				if stack[i] == ']' {
					//修改stack数据
					stack = stack[:i]
					break
				}
				tmp = append(tmp, stack[i])
			}
			for i := len(tmp) - 1; i >= 0; i-- {
				tmpStack = append(tmpStack, tmp[i])
			}
		} else if s[i] >= '0' && s[i] <= '9' {
			//数字，此时tmpStack内一定有字符串
			numStr := make([]byte, 0)
			var numTmp string
			var num int
			for i >= 0 {
				if s[i] < '0' || s[i] > '9' {
					break
				}
				numStr = append(numStr, s[i])
				i--
			}
			i++
			for i := len(numStr) - 1; i >= 0; i-- {
				numTmp += string(numStr[i])
			}
			num, _ = strconv.Atoi(numTmp)
			for i := 0; i < num; i++ {
				stack = append(stack, tmpStack...)
			}
			//清空tmpStack
			tmpStack = make([]byte, 0)
		} else {
			stack = append(stack, s[i])
		}
	}
	//翻转stack
	for i := len(stack) - 1; i >= 0; i-- {
		res = append(res, stack[i])
	}
	return string(res)
}
