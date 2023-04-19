package leetcode

/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func isHappy(n int) bool {
	record := map[int]bool{}
	for {
		sum := getSum(n)
		if sum == 1 {
			return true
		}

		if record[sum] {
			return false
		} else {
			record[sum] = true
		}
		n = sum
	}
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}

// @lc code=end
