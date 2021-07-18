/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g []int, s []int) (ans int) {
	sort.Ints(g)
	sort.Ints(s)
	m, n := len(g), len(s)
	for i, j := 0, 0; i < m && j < n; i++ {
		for j < n && g[i] > s[j] {
			j++
		}
		if j < n {
			ans ++
			j ++
		}
	}
	return
}
// @lc code=end

