/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func isHappy(n int) bool {
	step := func(n int) int{
		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n = n / 10
		}
		return sum
	}
	slow, fast := n, step(n)
	for slow != fast && fast != 1 {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}
// @lc code=end
