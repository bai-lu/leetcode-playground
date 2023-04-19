package leetcode

/*
 * @lc app=leetcode.cn id=707 lang=golang
 *
 * [707] 设计链表
 */

// @lc code=start
// type MyLinkedList struct {
// 	Val  int
// 	Next *MyLinkedList
// }

// func Constructor() MyLinkedList {
// 	return MyLinkedList{}
// }

// func (this *MyLinkedList) Get(index int) int {
// 	p := this.Next
// 	for i := 0; i < index; i++ {
// 		if p == nil {
// 			return -1
// 		}
// 		p = p.Next
// 	}
// 	if p == nil {
// 		return -1
// 	}
// 	return p.Val
// }

// func (this *MyLinkedList) AddAtHead(val int) {
// 	tmp := &MyLinkedList{Val: val}
// 	tmp.Next = this.Next
// 	this.Next = tmp
// }

// func (this *MyLinkedList) AddAtTail(val int) {
// 	p := this
// 	for p.Next != nil {
// 		p = p.Next
// 	}
// 	p.Next = &MyLinkedList{Val: val}
// }

// func (this *MyLinkedList) AddAtIndex(index int, val int) {
// 	p := this
// 	for i := 0; i < index; i++ {
// 		if p == nil {
// 			return
// 		}
// 		p = p.Next
// 	}
// 	if p == nil {
// 		return
// 	}
// 	tmp := &MyLinkedList{Val: val}
// 	tmp.Next = p.Next
// 	p.Next = tmp
// }

// func (this *MyLinkedList) DeleteAtIndex(index int) {
// 	p := this
// 	for i := 0; i < index; i++ {
// 		if p == nil {
// 			return
// 		}
// 		p = p.Next
// 	}
// 	if p.Next == nil {
// 		return
// 	}
// 	p.Next = p.Next.Next

// }

type MyLinkedList struct {
	head *ListNode
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{&ListNode{}, 0}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.size {
		return -1
	}
	cur := l.head
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (l *MyLinkedList) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *MyLinkedList) AddAtTail(val int) {
	l.AddAtIndex(l.size, val)
}

func (l *MyLinkedList) AddAtIndex(index, val int) {
	if index > l.size {
		return
	}
	index = max(index, 0)
	l.size++
	pred := l.head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	toAdd := &ListNode{val, pred.Next}
	pred.Next = toAdd
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.size {
		return
	}
	l.size--
	pred := l.head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	pred.Next = pred.Next.Next
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
// @lc code=end
