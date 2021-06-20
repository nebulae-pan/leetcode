/*
 * @lc app=leetcode.cn id=1002 lang=golang
 *
 * [1002] 查找常用字符
 */

// @lc code=start
func commonChars(words []string) (res []string) {
	minFreq := [26]int{}

	for i := range minFreq {
		minFreq[i] = math.MaxInt64
	}

	for _, word := range words {
		freq := [26]int{}
		for _, b := range word {
			freq[b - 'a']++
		}
		for i, v := range freq {
			if v < minFreq[i] {
				minFreq[i] = v
			}
		}
	}

	for i := 0; i < 26; i++ {
		for j := 0 ; j < minFreq[i]; j++ {
			res = append(res, string(i + 'a'))
		}
	}
	return
}
// @lc code=end

