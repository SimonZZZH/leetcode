//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
//注意：给定 n 是一个正整数。
//
//示例 1：
//
//输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/climbing-stairs
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package main

import "fmt"

func main() {
	//start := time.Now().UnixNano()
	//fmt.Println(climbStairs(45))
	//fmt.Println((time.Now().UnixNano() - start) / 1000000)
	//start = time.Now().UnixNano()
	//fmt.Println(climbStairs2(45))
	//fmt.Println((time.Now().UnixNano() - start) / 1000000)
	fmt.Println(climbStairs2(45))
}

//循环效率比递归高很多
//LeetCode中45递归超时

//循环
func climbStairs(n int) int {
	res := make([]int, n+1)
	res[0] = 1
	res[1] = 2
	if n == 1 {
		return 1
	}
	for x := 2; x < n; x++ {
		res[x] = res[x-1] + res[x-2]
	}
	return res[n-1]
}

//递归
func climbStairs2(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs(n-1) + climbStairs(n-2)
}
