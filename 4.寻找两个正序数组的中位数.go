package leetcode

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := []int{}
	i, j := 0, 0
	for i < len(nums1) || j < len(nums2) {
		if i == len(nums1) {
			merged = append(merged, nums2[j])
			j++
			continue
		}

		if j == len(nums2) {
			merged = append(merged, nums1[i])
			i++
			continue
		}
		if nums1[i] <= nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}
	if len(merged)%2 == 0 {
		left := len(merged)/2 - 1
		right := len(merged) / 2
		return float64(merged[left]+merged[right]) / 2

	} else {
		mid := len(merged) / 2
		return float64(merged[mid])
	}

}

// @lc code=end
