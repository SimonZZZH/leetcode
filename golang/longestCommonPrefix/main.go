package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"aa", "a"}))
}

//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minLength := len(strs[0])
	index := -1
	for x := 0; x < len(strs); x++ {
		if len(strs[x]) < minLength {
			minLength = len(strs[x])
		}

	}
	for x := 0; x < minLength; x++ {
		if comp(strs, x) {
			index = x
		} else {
			break
		}
	}
	if index == -1 {
		return ""
	}
	return strs[0][:index+1]
}
func comp(strs []string, index int) bool {
	count := 0
	for x := 0; x < len(strs); x++ {
		if strs[x][index] == strs[0][index] {
			count++
		} else {
			break
		}
	}
	if count == len(strs) {
		return true
	} else {
		return false

	}
}
