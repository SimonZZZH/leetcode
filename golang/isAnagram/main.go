//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
//示例 1:
//
//输入: s = "anagram", t = "nagaram"
//输出: true
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-anagram
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

import "fmt"

func main() {
	fmt.Println(isAnagram("asdfghjkl", "gasfdhjkl"))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sign := make([]int, 26)
	for _, v := range s {
		sign[v-'a']++
	}
	for _, v := range t {
		sign[v-'a']--
	}
	for _, v := range sign {
		if v != 0 {
			return false
		}
	}
	return true
}
