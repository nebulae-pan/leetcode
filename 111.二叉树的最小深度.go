/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	var dfs func(root *TreeNode, depth int) int
	dfs = func(node *TreeNode, depth int) int {
		if node.Left == nil && node.Right == nil{
			return depth
		}
		minValue := math.MaxInt64
		if node.Left != nil {
			minValue = min(minValue, dfs(node.Left, depth))
		}
		if node.Right != nil {
			minValue = min(minValue, dfs(node.Right, depth))
		}
		return minValue + 1
	}
	return dfs(root, 1)
}
// @lc code=end

