package leetcode

/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	numCount := map[int]int{}
	result := []int{}
	for _, v := range nums {
		numCount[v]++
	}
	for i := 0; i < k; i++ {
		index, max := 0, 0 // find top 1 frequent element's frequent
		for k, v := range numCount {
			if v > max {
				max = v
				index = k
			}
		}
		result = append(result, index)
		delete(numCount, index)
	}
	return result
}

// @lc code=end
