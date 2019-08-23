package main

import "math"

//二叉树的最大深度
func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//迭代 todo
//func maxDepth2(root *TreeNode) int {
//
//}

//递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		leftDepth := maxDepth(root.Left)
		rightDepth := maxDepth(root.Right)
		return int(math.Max(float64(leftDepth), float64(rightDepth)) + 1)
	}
}
