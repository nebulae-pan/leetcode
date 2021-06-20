/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left , right := 0, 0
	for leftN := root; leftN != nil; {
		left ++
		leftN = leftN.Left
	}
	for rightN := root; rightN != nil; {
		right ++
		rightN = rightN.Right
	}
	if left == right {
		return 1 << left - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}
// @lc code=end

