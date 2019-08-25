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

//给定一个二叉树，返回它的 前序 遍历。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	res = append(res, root.Val)
	if root.Left != nil {
		res = append(res, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		res = append(res, preorderTraversal(root.Right)...)
	}
	return res
}

//TODO 迭代
