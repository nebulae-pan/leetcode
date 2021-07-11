/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根节点到叶节点数字之和
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
func sumNumbers(root *TreeNode) (ans int) {
	var dfs func(root *TreeNode, preValue int) int
	dfs = func(root *TreeNode, preValue int) int{
		if root == nil {
			return 0
		}
		temp := preValue * 10 + root.Val
		if root.Left == nil && root.Right == nil {
			return temp
		}
		return dfs(root.Left, temp) + dfs(root.Right,temp)
	}
	
	return dfs(root, 0)
}
// @lc code=end

