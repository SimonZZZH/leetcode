package main

import (
	"math"
)

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

//判断是否为平衡二叉树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	//fmt.Println(maxDepth(root.Left), maxDepth(root.Right))
	//if math.Abs(float64(maxDepth(root.Left)-maxDepth(root.Right))) <= 1 {
	//	return true
	//}
	//递归，每个子树也是平衡二叉树
	if math.Abs(float64(maxDepth(root.Left)-maxDepth(root.Right))) <= 1 && isBalanced(root.Left) && isBalanced(root.Right) {
		return true
	}
	return false
}

//获取树最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		leftDepth := maxDepth(root.Left)
		rightDepth := maxDepth(root.Right)
		return int(math.Max(float64(leftDepth), float64(rightDepth)) + 1)
	}
}
