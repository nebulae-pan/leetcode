/*
 * @lc app=leetcode.cn id=525 lang=golang
 *
 * [525] 连续数组
 */

// @lc code=start
func findMaxLength(nums []int) (ans int) {
	mp := make(map[int]int)
	mp[0] = -1
	pre := 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i, v := range(nums) {
		if v == 0 {
			pre--
		} else {
			pre++
		}
		if count, ok := mp[pre]; ok {
			ans = max(ans, i - count)
		} else {
			mp[pre] = i
		}
	}
	return ans
}
// @lc code=end

