/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	len := len(nums) 
	for i := 0; i < len; i++ {
		for nums[i] > 0 && nums[i] < len && nums[i] != nums[nums[i] - 1] {
			nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]
		}
	}
	for i, v := range nums {
		if i + 1 != v {
			return i + 1
		}
	}
	return len + 1
}
// @lc code=end

