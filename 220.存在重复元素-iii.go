/*
 * @lc app=leetcode.cn id=220 lang=golang
 *
 * [220] 存在重复元素 III
 */

// @lc code=start
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	mp := make(map[int]int)
	abs := func(n int) int {
		if n < 0 {
			return n * -1
		}
		return n
	}
	bucketId := func (n, t int) int {
		if n >= 0 {
			return n / t
		}
		return (n + 1) / t - 1
	}
	for i, v := range nums {
		id := bucketId(v, t + 1) 
			
		if _, has := mp[id]; has {
			return true
		}
		if y, has := mp[id - 1]; has && abs(v - y) <= t {
			return true
		}
		if y, has := mp[id + 1]; has && abs(v - y) <= t {
			return true
		}
		mp[id] = v
		if i >= k {
			delete(mp, bucketId(nums[i - k], t + 1))
		}
	}
	return false
}
// @lc code=end

