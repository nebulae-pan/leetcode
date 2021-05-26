/*
 * @lc app=leetcode.cn id=1859 lang=golang
 *
 * [1859] 将句子排序
 */

// @lc code=start
import(
	"strings"
)
func sortSentence(s string) string {
	strs := strings.Split(s, " ")
	res := make([]string, len(strs))
	for i := 0; i < len(strs); i ++ {
		len := len(strs[i])
		res[strs[i][len - 1] - '1'] = strs[i][:len - 1]
	}
	return strings.Join(res, " ")
}
// @lc code=end

