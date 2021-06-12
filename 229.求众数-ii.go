/*
 * @lc app=leetcode.cn id=229 lang=golang
 *
 * [229] 求众数 II
 */

// @lc code=start
func majorityElement(nums []int) (res []int) {
	res = make([]int , 0)
	if nums == nil || len(nums) == 0 {
		return
	}
	num1, num2 := nums[0], nums[0]
	count1, count2 := 0, 0
	for _, v := range nums {
		if v == num1 {
			count1 ++
			continue
		}
		if v == num2 {
			count2 ++
			continue
		}
		if count1 == 0 {
			num1 = v
			count1++
			continue
		}
		if count2 == 0 {
			num2 = v
			count2++
			continue
		}
		count1--
		count2--
	}
	count1 = 0
	count2 = 0
	for _,v := range nums {
		if v == num1 {
			count1++
			continue
		}
		if v == num2 {
			count2++
		}
	}
	if count1 > len(nums) / 3 {
		res = append(res, num1)
	}
	if count2 > len(nums) / 3 {
		res = append(res, num2)
	}
	return
}
// @lc code=end

