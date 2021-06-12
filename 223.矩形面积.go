/*
 * @lc app=leetcode.cn id=223 lang=golang
 *
 * [223] 矩形面积
 */

// @lc code=start
func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	width1 := ax2 - ax1
	height1 := ay2 - ay1

	area1 := width1 * height1

	width2 := bx2 - bx1
	heigh2 := by2 - by1

	area2 := width2 * heigh2

	min := func (a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	across := func (start1, end1, start2, end2 int) int{
		if start1 < start2 {
			if start2 < end1 {
				return min(end1, end2) - start2
			}
			return 0
		}
		if start2 < start1 {
			if start1 < end2 {
				return min(end1, end2) - start1
			}
			return 0
		}
		return min(end2, end1) - start1
	}

	acrossWidth := across(ax1, ax2, bx1, bx2)
	acrossHeight := across(ay1, ay2, by1, by2)

	acrossArea := acrossWidth * acrossHeight
	return area1 + area2 - acrossArea
}
// @lc code=end

