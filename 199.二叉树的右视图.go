package leetcode

/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func rightSideView(root *TreeNode) []int {
	record := [][]int{}
	result := []int{}
	var levelOrder func(root *TreeNode, depth int)
	levelOrder = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if depth == len(record) {
			record = append(record, []int{})
		}
		record[depth] = append(record[depth], root.Val)
		levelOrder(root.Left, depth+1)
		levelOrder(root.Right, depth+1)
	}
	levelOrder(root, 0)
	for _, v := range record {
		result = append(result, v[len(v)-1])

	}
	return result
}

// @lc code=end
