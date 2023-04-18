package leetcode

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	// if n == 1 {
	// 	return 1
	// }
	// if n == 2 {
	// 	return 2
	// }
	// return climbStairs(n-1) + climbStairs(n-2)

	record := make([]int, 45)
	record[0] = 1
	record[1] = 2
	for i := 2; i < n; i++ {
		record[i] = record[i-1] + record[i-2]
	}
	return record[n-1]
}

// @lc code=end
