/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] 找不同
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	mp := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		v := s[i]
		if times, has := mp[v]; has {
			mp[v] = times + 1
		} else {
			mp[v] = 1
		}
	}
	for i := 0; i < len(t); i++ {
		v := t[i]
		if times, has := mp[v]; has {
			if times == 0 {
				return v
			}
			mp[v] = times - 1
		} else {
			return v
		}
	}
	return 'a'
}
// @lc code=end

