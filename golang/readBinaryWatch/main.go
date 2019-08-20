//二进制手表顶部有 4 个 LED 代表小时（0-11），底部的 6 个 LED 代表分钟（0-59）。
//
//每个 LED 代表一个 0 或 1，最低位在右侧。
//
//例如，上面的二进制手表读取 “3:25”。
//
//给定一个非负整数 n 代表当前 LED 亮着的数量，返回所有可能的时间。
//
//案例:
//
//输入: n = 1
//返回: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]
//
//
//
//注意事项:
//
//    输出的顺序没有要求。
//    小时不会以零开头，比如 “01:00” 是不允许的，应为 “1:00”。
//    分钟必须由两位数组成，可能会以零开头，比如 “10:2” 是无效的，应为 “10:02”。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-watch
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package main

import "fmt"

func main() {
	fmt.Println(readBinaryWatch(1))
}
func readBinaryWatch(num int) []string {
	//穷举。。。
	res := make([]string, 0)
	for i := 0; i < 12; i++ {
		sum_h := count1(i)
		for k := 0; k < 60; k++ {
			sum_m := count1(k)
			if sum_h+sum_m == num {
				tmp := fmt.Sprintf("%d:%02d", i, k)
				res = append(res, tmp)
			}
		}
	}
	return res
}

//统计二进制中1的个数
func count1(num int) int {
	var sum int
	for num != 0 {
		if num&1 == 1 {
			sum++
		}
		num = num >> 1
	}
	return sum
}
