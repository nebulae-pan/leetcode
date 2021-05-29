/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */

// @lc code=start
import(
	"fmt"
)
func search(nums []int, target int) bool {
	low , high := 0, len(nums) - 1

	for low <= high {
		mid := (low + high) >> 1
		v := nums[mid]
		// fmt.Printf("low:%d, high: %d, mid: %d\n", low, high, mid)

		if v == target {
			return true
		}
		if v == nums[low] && v == nums[high] {
			low++ 
			high--
			continue
		}
		if nums[low] > v {
			if target > v && target <= nums[len(nums) - 1] {
				low = mid + 1
			} else {
				high = mid - 1
			}

		} else {
			if target < v && target >= nums[low] {
				high = mid - 1
 			} else {
				low = mid + 1
			}
		}
	}
	return false
}
// @lc code=end

