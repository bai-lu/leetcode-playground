package leetcode

/*
 * @lc app=leetcode.cn id=1047 lang=golang
 *
 * [1047] 删除字符串中的所有相邻重复项
 */

// @lc code=start
func removeDuplicates(s string) string {
	stack := []rune{}
	// var top rune
	for _, c := range s {
		if len(stack) > 0 && stack[len(stack)-1] == c {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
		// if c != top {
		// 	stack = append(stack, c)
		// } else if len(stack) > 0 {
		// 	stack = stack[:len(stack)-1]
		// 	top = stack[len(stack)-1]
		// }

	}
	return string(stack)

}

// @lc code=end
