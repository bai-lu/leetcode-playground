package leetcode

/*
 * @lc app=leetcode.cn id=844 lang=golang
 *
 * [844] 比较含退格的字符串
 */

// @lc code=start
func backspaceCompare(s string, t string) bool {
	stack1 := []byte{}
	stack2 := []byte{}
	for _, v := range s {
		if v != '#' {
			stack1 = append(stack1, byte(v))
		} else if len(stack1) > 0 {
			stack1 = stack1[:len(stack1)-1]
		}
	}

	for _, v := range t {
		if v != '#' {
			stack2 = append(stack2, byte(v))
		} else if len(stack2) > 0 {
			stack2 = stack2[:len(stack2)-1]
		}
	}

	if len(stack1) == len(stack2) {
		for i := 0; i < len(stack1); i++ {
			if stack1[i] != stack2[i] {
				return false
			}
		}
		return true
	}
	return false

}

// @lc code=end
