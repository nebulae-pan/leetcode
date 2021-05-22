/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	low, high := 0, len(nums) - 1
	for low <= high {
		mid := (low + high) >> 1
		value := nums[mid]
		if value == target {
			first := mid
			for first > 0 && nums[first] == nums[first - 1] {
				first--
			}
			last := mid 
			for last < len(nums) - 1 && nums[last] == nums[last + 1] {
				last ++
			}

			return []int{first, last}
		}
		if value < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return []int{-1, -1}
}
// @lc code=end

