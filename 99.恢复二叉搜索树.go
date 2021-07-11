/*
 * @lc app=leetcode.cn id=99 lang=golang
 *
 * [99] 恢复二叉搜索树
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
func recoverTree(root *TreeNode)  {
	var t1, t2, pre *TreeNode
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if pre != nil && pre.Val > root.Val {
			if t1 == nil {
				t1 = pre
			}
			t2 = root
		}
		pre = root
		dfs(root.Right)
	}
	dfs(root)
	t1.Val, t2.Val = t2.Val, t1.Val
}
// @lc code=end

