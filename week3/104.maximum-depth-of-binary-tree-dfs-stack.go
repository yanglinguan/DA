/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
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
// dfs using stack
func maxDepth(root *TreeNode) int {
	if root == nil { return 0 }
	stack := make([]levelNode, 0)
	stack = append(stack, levelNode{1, root})
	level := 0
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top.level > level {
			level = top.level
		} 
		if top.node.Right != nil {
			stack = append(stack, levelNode{top.level+1, top.node.Right})
		}
		if top.node.Left != nil {
			stack = append(stack, levelNode{top.level+1, top.node.Left})
		}
	}
	return level
}

type levelNode struct {
	level int
	node *TreeNode
}
// @lc code=end

