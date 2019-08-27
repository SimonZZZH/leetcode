//给定两个字符串 A 和 B, 寻找重复叠加字符串A的最小次数，使得字符串B成为叠加后的字符串A的子串，如果不存在则返回 -1。
//
//举个例子，A = "abcd"，B = "cdabcdab"。
//
//答案为 3， 因为 A 重复叠加三遍后为 “abcdabcdabcd”，此时 B 是其子串；A 重复叠加两遍后为"abcdabcd"，B 并不是其子串。
//
//注意:
//
// A 与 B 字符串的长度在1和10000区间范围内。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/repeated-string-match
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(repeatedStringMatch("aaaaaaaaaab", "ba"))
	fmt.Println(repeatedStringMatch("abcd", "cdabcdab"))
	fmt.Println(repeatedStringMatch("abcd", "abcdabcdabcd"))
	fmt.Println(repeatedStringMatch("abcd", "cdabcdabcdabcdab"))
	fmt.Println(repeatedStringMatch("abcd", "abcd"))
}

//1、A长度大于等于B长度，
//则A或者A重复1次即可
//2、A比B短
//则A一定在B中循环，然后在首尾各加一个A则可覆盖B

func repeatedStringMatch(A string, B string) int {
	if len(A) < len(B) {
		var tmp string
		//最小长度B.length/A.length
		for i := 0; i < len(B)/len(A); i++ {
			tmp += A
		}
		if strings.Index(tmp, B) >= 0 {
			return len(B) / len(A)
		}
		//最多再在首尾各加一个
		for i := 0; i < 2; i++ {
			tmp += A
			if strings.Index(tmp, B) >= 0 {
				return len(B)/len(A) + i + 1
			}
		}

	} else {
		tmp := A + A
		if strings.Index(A, B) >= 0 {
			return 1
		}
		if strings.Index(tmp, B) >= 0 {
			return 2
		}

	}
	return -1
}
