package leetcode

/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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

// recursion
func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		result = append(result, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return result
}

// interation todo
// @lc code=end
