package leetcode

import (
	"fmt"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
	// res := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	res := maxSlidingWindow([]int{1, -1}, 1)
	fmt.Println(res)

}
