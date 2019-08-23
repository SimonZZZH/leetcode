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
//给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return nil
	}
	getSlice(root, 0, &res)
	for x, y := 0, len(res)-1; x < y; x, y = x+1, y-1 {
		res[x], res[y] = res[y], res[x]
	}
	return res
}

func getSlice(node *TreeNode, dep int, res *[][]int) {
	if node == nil {
		return
	}
	if dep == len(*res) {
		*res = append(*res, []int{node.Val})
	} else {
		(*res)[dep] = append((*res)[dep], node.Val)
	}
	getSlice(node.Left, dep+1, res)
	getSlice(node.Right, dep+1, res)
}
