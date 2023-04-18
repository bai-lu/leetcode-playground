package leetcode

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	left := 0
	right := x / 2

	if x == 1 {
		return 1
	}

	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}

/*
note:
mid := left + (right-left)/2
mid * mid == x
return right
*/

// @lc code=end
