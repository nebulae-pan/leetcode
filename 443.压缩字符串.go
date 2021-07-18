/*
 * @lc app=leetcode.cn id=443 lang=golang
 *
 * [443] 压缩字符串
 */

// @lc code=start
import(
	"strconv"
)
func compress(chars []byte) (ans int) {
	if len(chars) == 1 {
		return 1
	}
	write := 0
	preIndex := 0
	for i, v := range(chars) {
		if  i == len(chars) - 1 || v != chars[i + 1] {
			ans++
			currNums := i - preIndex + 1
			preIndex = i + 1
			chars[write] = v
			write++
			if currNums != 1 {
				numStr := strconv.Itoa(currNums)
				for _, nv := range([]byte(numStr)) {
					chars[write] = nv
					write++
					ans++
				}
			}
			currNums = 1
		}
	}
	return
}
// @lc code=end

