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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//给定一个二叉树，返回它的中序 遍历。
//递归
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	if root.Left != nil {
		res = append(res, inorderTraversal(root.Left)...)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		res = append(res, inorderTraversal(root.Right)...)
	}
	return res
}

//迭代
//TODO
