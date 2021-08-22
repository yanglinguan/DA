/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
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
// dfs stack
func minDepth(root *TreeNode) int {
	if root == nil { return 0 }
	stack := make([]levelNode, 0)
	stack = append(stack, levelNode{1, root})
	level := math.MaxInt32

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top.node.Left == nil && top.node.Right == nil && top.level < level {
			level = top.level
		}
		if top.node.Left != nil {
			stack = append(stack, levelNode{top.level+1, top.node.Left})
		}
		if top.node.Right != nil {
			stack = append(stack, levelNode{top.level+1, top.node.Right})
		}
	}
	return level
}

type levelNode struct {
	level int
	node *TreeNode
}
// @lc code=end

