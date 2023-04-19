package leetcode

import (
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	src := []int{1, 1, 2}
	dst := []int{3, 3, 4, 4}
	copy(dst, src)
	fmt.Println(src)
	fmt.Println(dst)
}

func TestFindK(t *testing.T) {
	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 3}
	node1.Next = &node2
	node2.Next = &node3
	head := &node1
	res := findK(head, 2)
	fmt.Println(res.Val)
}

func findK(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	curr := dummyHead
	for i := 1; i <= k; i++ {
		curr = curr.Next
	}
	return curr
}
