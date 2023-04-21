package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}
