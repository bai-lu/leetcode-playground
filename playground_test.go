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
