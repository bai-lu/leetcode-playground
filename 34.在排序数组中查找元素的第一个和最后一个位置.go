package leetcode

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	mark := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			mark = mid
			break
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// target not found
	if mark == -1 {
		return []int{-1, -1}
	}

	// target found
	// search right border
	for i := mark; i < len(nums); i++ {
		if nums[i] == nums[mark] {
			right = i
		} else {
			break
		}
	}

	// search left border

	for i := mark; i >= 0; i-- {
		if nums[i] == nums[mark] {
			left = i
		} else {
			break
		}
	}
	return []int{left, right}

}

/*
1. 分别查找左右边界
2. 查找到左边界, 向左右扩展
*/

// @lc code=end
