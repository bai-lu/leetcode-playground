package leetcode

/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	res := []int{}
	for _, v := range nums {
		m[v]++
	}

	for i := 0; i < k; i++ {
		top := 0
		topKey := 0
		for k, v := range m {
			if v > top {
				top = v
				topKey = k
			}
		}
		res = append(res, topKey)
		delete(m, topKey)
	}
	return res
}

// @lc code=end
