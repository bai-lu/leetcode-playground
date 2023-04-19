package leetcode

/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	record := map[byte]int{}
	for _, v := range s {
		record[byte(v)]++
	}
	for _, v := range t {
		_, ok := record[byte(v)]
		if !ok {
			return false
		}
		record[byte(v)]--
	}

	for _, v := range record {
		if v != 0 {
			return false
		}
	}
	return true
}

/*
note: hash table
如果包含unicode, 使用rune类型
*/

// @lc code=end
