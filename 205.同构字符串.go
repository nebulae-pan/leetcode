/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	mp1 := make(map[byte]byte)
	mp2 := make(map[byte]byte)

	for i := range(s) {
		x := s[i]
		y := t[i]
		v1, ok1 := mp1[x]
		v2, ok2 := mp2[y]
		if ok1 && v1 != y || ok2 && v2 != x {
			return false
		}
		mp1[x] = y
		mp2[y] = x
	}
	return true
}
// @lc code=end

