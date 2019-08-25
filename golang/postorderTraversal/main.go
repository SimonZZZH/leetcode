package main

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

//给定一个二叉树，返回它的 后序遍历。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	if root.Left != nil {
		res = append(res, postorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		res = append(res, postorderTraversal(root.Right)...)
	}
	res = append(res, root.Val)
	return res
}

//TODO 迭代
