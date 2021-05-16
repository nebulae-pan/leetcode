/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) (result [][]int) {
	combination := make([]int, 0)
	slen := len(candidates)
	var dfs func(target, index int)
	dfs = func(target, index int) {
		if target == 0 {
			result = append(result, append([]int{}, combination...))
			return
		}
		for i := index; i < slen; i++ {
			v := candidates[i]
			if i > index && candidates[i - 1] == v {
				continue
			}
			if target - v >= 0 {
				combination = append(combination, v)
				dfs(target - v, i + 1)
				combination = combination[:len(combination) - 1]
			} else {
				break
			}
		}
	}
	sort.Ints(candidates)
	dfs(target, 0)
	return result
}
// @lc code=end

