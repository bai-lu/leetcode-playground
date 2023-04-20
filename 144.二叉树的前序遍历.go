package leetcode

/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
// func preorderTraversal(root *TreeNode) []int {
// 	result := []int{}
// 	var preOrder func(root *TreeNode)
// 	preOrder = func(root *TreeNode) {
// 		if root == nil {
// 			return
// 		}
// 		result = append(result, root.Val)
// 		preOrder(root.Left)
// 		preOrder(root.Right)
// 	}
// 	preOrder(root)
// 	return result
// }

// iteration
func preorderTraversal(root *TreeNode) []int {
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

		if tmp.Right != nil {
			stack = append(stack, tmp.Right)
		}
		if tmp.Left != nil {
			stack = append(stack, tmp.Left)
		}
	}
	return result
}

/*
note
1. 递归 recursion
2. 迭代  iteration
*/

// @lc code=end
