package leetcode

/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	result := []int{}
	record := map[int]struct{}{}
	for _, v := range nums1 {
		record[v] = struct{}{}
	}
	for _, v := range nums2 {
		if _, ok := record[v]; ok {
			result = append(result, v)
			delete(record, v)
		}
	}
	return result
}

// @lc code=end
