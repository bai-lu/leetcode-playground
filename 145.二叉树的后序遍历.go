package leetcode

/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
// func postorderTraversal(root *TreeNode) []int {
// 	result := []int{}
// 	var postOrder func(root *TreeNode)
// 	postOrder = func(root *TreeNode) {
// 		if root == nil {
// 			return
// 		}
// 		postOrder(root.Left)
// 		postOrder(root.Right)
// 		result = append(result, root.Val)
// 	}
// 	postOrder(root)
// 	return result
// }

// iteration
func postorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	result := []int{}
	if root == nil {
		return result
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, tmp.Val)
		if tmp.Left != nil {
			stack = append(stack, tmp.Left)
		}
		if tmp.Right != nil {
			stack = append(stack, tmp.Right)
		}
	}
	// reverse
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
	}
	return result
}

// @lc code=end
