package leetcode

/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)
	if root.Left == nil && root.Right != nil {
		return rightDepth + 1
	}
	if root.Left != nil && root.Right == nil {
		return leftDepth + 1
	}
	return min(leftDepth, rightDepth) + 1

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
