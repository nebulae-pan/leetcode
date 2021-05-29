/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除有序数组中的重复项 II
 */

// @lc code=start
import(
	"fmt"
)
func removeDuplicates(nums []int) int {
	len := len(nums)
	if len <= 2 {
		return len
	}
	index := 2
	for i := 2; i < len ; i ++ {
		if nums[i] != nums[index - 2] {
			nums[index] = nums[i]
			index ++
		}
	}
	return index
}
// @lc code=end

