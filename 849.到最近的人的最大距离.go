/*
 * @lc app=leetcode.cn id=849 lang=golang
 *
 * [849] 到最近的人的最大距离
 */

// @lc code=start
func maxDistToClosest(seats []int) (ans int) {
	len := len(seats)
	prev, future := -1, 0
	ans = 0
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i, v := range(seats) {
		if v == 1 {
			prev = i
		} else {
			for future < len && seats[future] == 0 || future < i {
				future++
			}
			left := 0
			if prev == -1 {
				left = len
			} else {
				left = i - prev
			}
			right := 0
			if future == len {
				right = len
			} else {
				right = future - i
			}
			ans = max(ans, min(right, left))
		}
	}
	return
}
// @lc code=end

