package leetcode

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if v, ok := m[target-nums[i]]; ok {
			return []int{v, i}
		} else {
			// m[target-i] = nums[i]
			m[nums[i]] = i
		}
	}
	return []int{0, 0}

}

// @lc code=end
