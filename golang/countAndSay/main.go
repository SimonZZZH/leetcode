//报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：
//
//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//
//1 被读作  "one 1"  ("一个一") , 即 11。
//11 被读作 "two 1s" ("两个一"）, 即 21。
//21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。
//
//给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。
//
//注意：整数顺序将表示为一个字符串。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/count-and-say
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(6), "===")
}

func countAndSay(n int) string {
	//res := make([]string, 0)
	//res[0] = "1"
	var res string
	if n == 1 {
		return "1"
	}
	var index int
	index = 1
	//var count int
	//count = 1
	res = countAndSay(n - 1)
	if len(res) == 1 {
		tmp := "1" + string(res[index-1])
		return tmp
	}
	var tmp string
	//fmt.Println(res, "_+_+_+")
	for index < len(res) {
		count := 1
		for x := index; x < len(res); x++ {
			if res[x-1] == res[x] {
				count++
				index++
				continue
			}
			break
		}
		tmp = tmp + strconv.Itoa(count) + string(res[index-1])
		//fmt.Println(tmp, index, "-----")
		index++
		count = 1
		if index == len(res) {
			tmp = tmp + strconv.Itoa(count) + string(res[index-1])
			//fmt.Println(tmp, "--+++---")
			index++
			count = 1
		}
	}
	return tmp
}
