package main

import "fmt"

func main() {
	matrix := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	fmt.Println(matrix)
	rotate(matrix)
	fmt.Println(matrix)

}

//给定一个 n × n 的二维矩阵表示一个图像。
//
//将图像顺时针旋转 90 度。
//
//说明：
//
//你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/rotate-image
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//转置+翻转
func rotate(matrix [][]int) {
	n := len(matrix)
	if n < 2 {
		return
	}
	//转置
	for row := 0; row < n; row++ {
		for column := row; column < n; column++ {
			matrix[row][column], matrix[column][row] = matrix[column][row], matrix[row][column]
		}
	}
	//翻转
	for row := 0; row < n; row++ {
		for column := 0; column < n/2; column++ {
			matrix[row][column], matrix[row][n-1-column] = matrix[row][n-1-column], matrix[row][column]
		}
	}
}
