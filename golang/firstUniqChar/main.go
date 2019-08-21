//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
//
//案例:
//
//s = "leetcode"
//返回 0.
//
//s = "loveleetcode",
//返回 2.
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/first-unique-character-in-a-string
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar(""))
}

//改进，使用数组保存，两次遍历,16ms,5.8mb
func firstUniqChar(s string) int {
	sign := make([]int, 26)
	for _, v := range s {
		sign[v-'a']++
	}
	for k, v := range s {
		if sign[v-'a'] == 1 {
			return k
		}
	}
	return -1
}

//map，96ms，5.9mb，使用range效率比循环高
func firstUniqChar2(s string) int {
	signMap := make(map[byte]int)
	for x := 0; x < len(s); x++ {
		if signMap[s[x]] > 0 {
			signMap[s[x]]++
		} else {
			signMap[s[x]] = 1
		}
	}
	for x := 0; x < len(s); x++ {
		if signMap[s[x]] == 1 {
			return x
		}
	}
	return -1
}
