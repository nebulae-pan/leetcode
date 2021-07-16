/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 */

// @lc code=start
func findDuplicates(nums []int) (ans []int) {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	for _, n := range(nums) {
		index := abs(n)
		i := nums[index - 1]
		if i < 0 {
			ans = append(ans, index)
		} else {
			nums[index - 1] = -nums[index - 1]
		}
	}
	return ans
}
// @lc code=end

