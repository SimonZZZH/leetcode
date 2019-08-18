package main

import "fmt"

func main() {
	//array := []int{2, 4, 2, 2, 2}
	//	//quickSort(array, 0, len(array)-1)
	//	//fmt.Println(array)
	fmt.Println(arrayPairSum([]int{1, 4, 2, 2, 2, 3, 4, 5, 3, 2}))
}

//给定长度为 2n 的数组,
// 你的任务是将这些数分成 n 对, 例如 (a1, b1), (a2, b2), ..., (an, bn) ，
// 使得从1 到 n 的 min(ai, bi) 总和最大。
//输入: [1,4,3,2]
//
//输出: 4
//解释: n 等于 2, 最大总和为 4 = min(1, 2) + min(3, 4).

func arrayPairSum(nums []int) int {
	quickSort(nums, 0, len(nums)-1)
	res := 0
	for x := 0; x < len(nums); x = x + 2 {
		res += nums[x]
	}
	return res
}

//先实现快排
func quickSort(array []int, start, end int) {
	if start == end {
		return
	}
	l := start
	r := end
	//据说能提高效率
	array[l], array[(l+r)/2] = array[(l+r)/2], array[l]
	//fmt.Println(array[start : end+1])
	tmp := array[l]
	for l < r {
		for l < r {
			if array[r] < tmp {
				array[l] = array[r]
				l++
				break
			} else {
				r--
			}
		}
		//fmt.Println(l, r, "==")
		if l == r {
			tmp, array[r] = array[r], tmp
			//fmt.Println(array, "--")
			quickSort(array, start, r)
			if r == end {
				break
			}
			quickSort(array, r+1, end)
			break
		}
		for l < r {
			if array[l] > tmp {
				array[r] = array[l]
				r--
				break
			} else {
				l++
			}
		}
		//fmt.Println(l, r, "===")
		if l == r {
			tmp, array[r] = array[r], tmp
			//fmt.Println(array, "++")
			quickSort(array, start, r)
			if r == end {
				break
			}
			quickSort(array, r+1, end)
			break
		}
	}
}
