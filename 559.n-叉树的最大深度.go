package leetcode

/*
 * @lc app=leetcode.cn id=559 lang=golang
 *
 * [559] N 叉树的最大深度
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	max := 0
	for i := 0; i < len(root.Children); i++ {
		childDepth := maxDepth(root.Children[i])
		if childDepth > max {
			max = childDepth
		}
	}
	return max + 1

}

// @lc code=end
