package leetcode

/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	slow := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[slow] = nums[i]
			slow++
		}
	}

	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}

}

// @lc code=end
