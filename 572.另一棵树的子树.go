package leetcode

/*
 * @lc app=leetcode.cn id=572 lang=golang
 *
 * [572] 另一棵树的子树
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
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var preOrder func(root *TreeNode, subRoot *TreeNode) bool
	preOrder = func(root *TreeNode, subRoot *TreeNode) bool {
		if root == nil {
			return false
		}
		if check(root, subRoot) {
			return true
		}
		return preOrder(root.Left, subRoot) || preOrder(root.Right, subRoot)
	}
	return preOrder(root, subRoot)

}

func check(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return check(p.Left, q.Left) && check(p.Right, q.Right)
}

/*
一次AC, 厉害了
*/
// @lc code=end
