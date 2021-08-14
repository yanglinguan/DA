/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
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
 // iteration
func preorderTraversal(root *TreeNode) []int {
  stack := make([]colorNode, 0)
	stack = append(stack, colorNode{0, root})
	res := []int{}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top.node == nil { continue }
		if top.color == 0 {
			top.color = 1 
			stack = append(stack, colorNode{0, top.node.Right})
			stack = append(stack, colorNode{0, top.node.Left})
			stack = append(stack, top)
		} else {
			res = append(res, top.node.Val)
		}
	}
	return res
}

type colorNode struct {
	color int
	node *TreeNode
}
// @lc code=end

