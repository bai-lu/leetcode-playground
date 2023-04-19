package leetcode

/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
// Time Limit Exceeded
func minSubArrayLen(target int, nums []int) int {
	// left, right := 0, 0
	minLength := len(nums) + 1
	left := 0

	// 子数组求和比较耗时
	// arrSum := func(nums []int) (sum int) {
	// 	for i := 0; i < len(nums); i++ {
	// 		sum += nums[i]
	// 	}
	// 	return
	// }

	sum := 0
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			if right-left+1 < minLength {
				minLength = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}
	if minLength > len(nums) {
		return 0
	}
	return minLength
}

/*
note: 滑动窗口
subLength = j - i + 1
滑动窗口求和通过上个窗口的结果计算, 而不是进行遍历
*/

// @lc code=end
