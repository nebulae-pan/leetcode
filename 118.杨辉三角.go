/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
import(
	"fmt"
)
func generate(numRows int) (ans [][]int) {
	ans = make([][]int, 0)
	for i := 1; i <= numRows; i++ {
		temp := make([]int, i)
		temp[0] = 1
		temp[i - 1] = 1
		for j := 1; j < i - 1; j++ {
			temp[j] = ans[i - 2][j - 1] + ans[i - 2][j]
		}
		ans = append(ans, temp)
	}
	return
}
// @lc code=end

