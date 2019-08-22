//给定两个字符串 s 和 t，判断它们是否是同构的。
//
//如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。
//
//所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。
//
//示例 1:
//
//输入: s = "egg", t = "add"
//输出: true
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/isomorphic-strings
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isIsomorphic("aannpp", "kkopgg"))
}

//每个字符在自身第一次出现的位置相同，妙啊。。。大佬的思维就是不一样
func isIsomorphic(s string, t string) bool {
	for i := 0; i < len(s); i++ {
		if strings.IndexByte(s, s[i]) != strings.IndexByte(t, t[i]) {
			return false
		}
	}
	return true
}
