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

//对称二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(L, R *TreeNode) bool {
	if L == nil && R == nil {
		return true
	}
	if L == nil || R == nil {
		return false
	}
	if L.Val == R.Val && isMirror(L.Left, R.Right) && isMirror(L.Right, R.Left) {
		return true
	}
	return false
}
