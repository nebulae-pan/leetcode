/*
 * @lc app=leetcode.cn id=865 lang=golang
 *
 * [865] 具有所有最深节点的最小子树
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
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left == right {
		return root
	}
	if left < right {
		return subtreeWithAllDeepest(root.Right)
	} else {
		return subtreeWithAllDeepest(root.Left)
	}
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left) , maxDepth(root.Right)) + 1
}

func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}
// @lc code=end

