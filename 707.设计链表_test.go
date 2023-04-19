package leetcode

import (
	"fmt"
	"testing"
)

func TestDesignListNode(t *testing.T) {
	l := Constructor()
	l.AddAtHead(1)
	l.AddAtHead(2)
	l.AddAtHead(3)
	// l.Get(2)
	fmt.Println(l)

}
