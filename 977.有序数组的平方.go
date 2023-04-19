package leetcode

/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */

// @lc code=start
func sortedSquares(nums []int) []int {
	result := []int{}
	left := 0
	right := len(nums) - 1
	for left <= right {
		leftSquares := nums[left] * nums[left]
		rightSquares := nums[right] * nums[right]
		if leftSquares < rightSquares {
			result = append(result, rightSquares)
			right--
		} else {
			result = append(result, leftSquares)
			left++
		}
	}
	// reverse result
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
	return result
}

/*
note:
step1: 双指针, 分别指向slice首尾, 向中心移动
step2: 反转result
*/

// @lc code=end
