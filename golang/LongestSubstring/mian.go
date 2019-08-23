//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//示例 1:
//
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("aabb"))
}

func lengthOfLongestSubstring(s string) int {
	var longest int
	sByte := []byte(s)
	//count := make([]int, 0)
	if len(s) <= 1 {
		return len(s)
	}
	var startIndex, endIndex int
	for x := 1; x < len(sByte); x++ {
		//fmt.Println("++", string(sByte[startIndex:x]), string(sByte[x]), "++")
		endIndex = x
		if bytes.ContainsRune(sByte[startIndex:x], rune(sByte[x])) {
			if longest < endIndex-startIndex {
				longest = endIndex - startIndex
			}
			//count = append(count, endIndex-startIndex)
			startIndex = bytes.LastIndexByte(sByte[:x], sByte[x]) + 1
			//fmt.Println(startIndex)
		}
	}
	if longest < endIndex-startIndex+1 {
		longest = endIndex - startIndex + 1
	}
	//count = append(count, endIndex-startIndex+1)
	//fmt.Println(count)
	//sort.Ints(count)
	//fmt.Println(count)
	//longest = count[len(count)-1]
	return longest
}
