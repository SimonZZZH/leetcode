//给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
//
//如果不存在最后一个单词，请返回 0 。
//
//说明：一个单词是指由字母组成，但不包含任何空格的字符串。
//
//示例:
//
//输入: "Hello World"
//输出: 5
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/length-of-last-word
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLastWord("asss"))
}
func lengthOfLastWord(s string) int {
	if len(s) < 1 {
		return 0
	}
	//if s[len(s)-1] == ' ' {
	//	return 0
	//}
	var r, l int
	l = -1
	for x := len(s) - 1; x >= 0; x-- {
		if s[x] != ' ' {
			r = x
			break
		}
	}
	for x := r; x >= 0; x-- {
		if s[x] == ' ' {
			l = x
			break
		}
	}
	fmt.Println(r, l)
	return r - l
}
