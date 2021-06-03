/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	dummy := &ListNode{0, head}
	pre := dummy
	anchor := dummy
	p := dummy
	for p.Next != nil {
		if p.Next.Val >= x {
			pre = p
			anchor = p.Next
			break
		}
		p = p.Next
	}
	if anchor == dummy {
		return dummy.Next
	}
	p = anchor
	for p.Next != nil {
		if p.Next.Val < x {
			temp := p.Next.Next
			pre.Next = p.Next
			p.Next.Next = anchor
			pre = p.Next
			p.Next = temp
		} else {
			p = p.Next
		}
	}
	return dummy.Next
}
// @lc code=end

