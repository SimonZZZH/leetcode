package main

import "fmt"

func main() {
	fmt.Println(countBits(6))
}

//给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。

//给出时间复杂度为O(n*sizeof(integer))的解答非常容易。但你可以在线性时间O(n)内用一趟扫描做到吗？
//要求算法的空间复杂度为O(n)。
//你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的 __builtin_popcount）来执行此操作。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/counting-bits
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//偶数调用函数计算汉明重量，比该偶数大一的奇数直接加1即可
func countBits(num int) []int {
	res := make([]int, num+1)
	for x := 0; x < num+1; x = x + 2 {
		res[x] = hammingWeight(uint32(x))
		if x+1 < num+1 {
			res[x+1] = res[x] + 1
		}
	}
	return res
}

func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		if num%2 != 0 {
			count++
		}
		num = num >> 1
	}
	return count
}
