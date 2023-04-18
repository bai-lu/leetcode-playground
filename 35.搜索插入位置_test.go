package leetcode

import "testing"

func Test_searchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 7
	result := searchInsert(nums, target)
	println(result)

}
