/*
 * @lc app=leetcode.cn id=526 lang=golang
 *
 * [526] 优美的排列
 */

// @lc code=start
func countArrangement(n int) (ans int) {
	visited := make([]bool, n + 1)
	var calculate func(pos int)
	calculate = func(pos int) {
		if pos > n {
			ans++
		}
		for i := 1; i <= n; i++ {
			if !visited[i] && (pos % i == 0 || i % pos == 0) {
				visited[i] = true
				calculate(pos + 1)
				visited[i] = false
			}
		}
	}
	calculate(1)
	return 
}
// @lc code=end

