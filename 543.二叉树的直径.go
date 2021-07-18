/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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


func diameterOfBinaryTree(root *TreeNode) int {
	max := func (a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxV := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		maxV = max(maxV, left + right + 1)
		return max(left, right) + 1
	}
	dfs(root)
	return maxV - 1
}
// @lc code=end

