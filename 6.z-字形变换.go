/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	len := len(s)
	step := 2 * numRows - 2
	var ss []byte
	for i := 0 ; i < numRows; i++ {
		for j := 0; j + i < len; j += step {
			ss = append(ss, s[j + i])
			if i != 0 && i != numRows - 1 && j + step - i < len {
				ss = append(ss, s[j + step - i])
			}
		}
	}
	return string(ss)
}
// @lc code=end

