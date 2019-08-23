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

//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
//假设一个二叉搜索树具有如下特征：
//
//节点的左子树只包含小于当前节点的数。
//节点的右子树只包含大于当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	resSlice := getSlice(root)
	for x := 1; x < len(resSlice); x++ {
		if resSlice[x] <= resSlice[x-1] {
			return false
		}
	}
	return true
}

func getSlice(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	if root.Left != nil {
		res = append(res, getSlice(root.Left)...)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		res = append(res, getSlice(root.Right)...)
	}
	return res
}

//todo
//迭代
