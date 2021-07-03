/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start
func getRow(rowIndex int) (ans []int) {
	ans = make([]int, rowIndex + 1)
	ans[0] = 1
	for i := 1 ; i <= rowIndex; i++ {
		ans[i] = ans[i - 1] * (rowIndex - i + 1) / i
	}
	return
}
// @lc code=end

