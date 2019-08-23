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

//给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
//递归
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	getSlice(root, 0, &res)
	return res
}

func getSlice(root *TreeNode, dep int, res *[][]int) {
	if root == nil {
		return
	}
	if dep == len(*res) {
		*res = append(*res, []int{root.Val})
	} else {
		(*res)[dep] = append((*res)[dep], root.Val)
	}
	getSlice(root.Left, dep+1, res)
	getSlice(root.Right, dep+1, res)
}
