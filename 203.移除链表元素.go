/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode {0, head}
	p := dummy
	node := dummy.Next
	for node != nil {
		if node.Val == val {
			temp := node.Next
			node.Next = nil
			p.Next = temp
			node = temp
		} else {
			p = p.Next
			if p != nil {
				node = p.Next
			} else {
				node = p
			}
		}
	}
	return dummy.Next
}
// @lc code=end

