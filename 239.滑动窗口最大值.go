package leetcode

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	result := []int{}
	for i := 0; i < len(nums)-k+1; i++ {
		subNums := nums[i : i+k]
		maxFunc := func(nums []int) int {
			max := math.MinInt
			for i := 0; i < len(nums); i++ {
				if nums[i] > max {
					max = nums[i]
				}
			}
			return max

		}
		result = append(result, maxFunc(subNums))
		fmt.Println(subNums)
	}
	return result

}

// @lc code=end
