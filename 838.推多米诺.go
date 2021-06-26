/*
 * @lc app=leetcode.cn id=838 lang=golang
 *
 * [838] 推多米诺
 */

// @lc code=start
func pushDominoes(dominoes string) string {
	len := len(dominoes)
	diff := make([]int, len)
	weight := 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i, v := range(dominoes) {
		if v == 'R' {
			weight = len
		} else if v == 'L' {
			weight = 0
		} else {
			weight = max(0, weight - 1)
		}
		diff[i] = weight
	}
	for i := len - 1 ; i >= 0; i-- {
		v := dominoes[i]
		if v == 'L' {
			weight = len
		} else if v == 'R' {
			weight = 0
		} else {
			weight = max(0, weight - 1)
		}
		diff[i] = diff[i] - weight
	}
	bytes := make([]byte, 0)
	for i := 0; i < len; i++ {
		if diff[i] == 0 {
			bytes = append(bytes, '.')
			continue
		}
		if diff[i] > 0 {
			bytes = append(bytes, 'R')
		} else {
			bytes = append(bytes, 'L')
		}
	}
	return string(bytes)
}
// @lc code=end

