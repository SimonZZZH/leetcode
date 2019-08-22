//给定一个二叉树，它的每个结点都存放着一个整数值。
//
//找出路径和等于给定数值的路径总数。
//
//路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
//
//二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。
//
//示例：
//
//root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/path-sum-iii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

func main() {

}

//Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//双层递归
func pathSum(root *TreeNode, sum int) int {

}

//递归，找出当前节点向下满足要求的路径
func pathSumForNode(node *TreeNode, sum int) int {
	var res int
	tmpSum := node.Val
}

//一次遍历
//TODO
