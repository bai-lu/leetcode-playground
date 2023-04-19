package leetcode

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	curr := dummyHead
	len := getListLength(head)
	for i := 0; i < len-n; i++ {
		curr = curr.Next

	}
	curr.Next = curr.Next.Next
	return dummyHead.Next
}

func getListLength(head *ListNode) int {
	count := 0
	for head != nil {
		head = head.Next
		count++
	}
	return count
}

// @lc code=end
