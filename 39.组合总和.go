/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) (result [][]int) {
	combination := make([]int, 0)
	var dfs func(target, index int)
	dfs = func(target, index int) {
		if len(candidates) == index {
			return
		}
		if target == 0 {
			result = append(result, append([]int{}, combination...))
			return
		}
		// 直接跳过
		dfs(target, index + 1)
		v := candidates[index]
		if target - v >= 0 {
			combination = append(combination, v)
			dfs(target - v, index)
			combination = combination[:len(combination) - 1]
		}
	}
	dfs(target, 0)
	return result
}
// @lc code=end

