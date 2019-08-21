package main

import (
	"fmt"
)

func main() {
	fmt.Println(generate(5))
}

//杨辉三角
func generate(numRows int) [][]int {
	res := [][]int{[]int{1}, []int{1, 1}}
	if numRows < 3 {
		return res[:numRows]
	}
	for x := 2; x < numRows; x++ {
		row := []int{1}
		for y := 0; y < len(res[x-1])-1; y++ {
			row = append(row, res[x-1][y]+res[x-1][y+1])
		}
		row = append(row, 1)
		res = append(res, row)
	}
	return res
}
