package main

import "math"

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
//给定一个二叉树，找出其最小深度。
//
//最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//
//说明: 叶子节点是指没有子节点的节点。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left != nil && root.Right != nil {
		return int(math.Min(float64(minDepth(root.Right)), float64(minDepth(root.Left))) + 1)
	}
	if root.Left != nil {
		return minDepth(root.Left) + 1
	} else {
		return minDepth(root.Right) + 1
	}
}

//func minDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	if root.Left != nil && root.Right != nil {
//		return int(math.Min(float64(getMinDepth(root.Right)), float64(getMinDepth(root.Left))) + 1)
//	}
//	if root.Left != nil {
//		return getMinDepth(root.Left) + 1
//	} else {
//		return getMinDepth(root.Right) + 1
//	}
//}
//
//func getMinDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	} else {
//		lDepth := minDepth(root.Left)
//		rDepth := minDepth(root.Right)
//		return int(math.Min(float64(lDepth), float64(rDepth)) + 1)
//	}
//}
