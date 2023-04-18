package leetcode

import (
	"fmt"
	"testing"
)

func Test_searchRange(t *testing.T) {
	nums := []int{5, 7, 8, 8, 8, 8, 8, 8, 9}
	target := 8
	searchRange(nums, target)

}

func Test_printStar(t *testing.T) {
	str := `*
**
***
****
*****
****
***
**
*`
	fmt.Println(str)
}
