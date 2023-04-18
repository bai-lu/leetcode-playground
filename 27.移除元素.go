package leetcode

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	slow := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[slow] = nums[i]
			slow++
		}
	}
	nums = nums[:slow]
	return slow

}

/*
明确 slow 指的是什么
*/
// @lc code=end
