package leetcode

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
// func removeDuplicates(nums []int) int {
// 	slow := 0
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] != nums[i-1] {
// 			nums[slow] = nums[i]
// 			slow++
// 		}
// 	}
// 	nums = nums[:slow]
// 	return len(nums)

func removeDuplicates(nums []int) int {
	m := map[int]struct{}{}
	result := []int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			continue
		} else {
			m[nums[i]] = struct{}{}
			result = append(result, nums[i])
		}
	}

	// for i := 0; i < len(result); i++ {
	// 	nums[i] = result[i]
	// }
	copy(nums, result)

	return len(result)
}

/*
note : 直接看是个hash table 题目, 这里的解法是hash table
但是可以用快慢指针的思路
*/
// @lc code=end
