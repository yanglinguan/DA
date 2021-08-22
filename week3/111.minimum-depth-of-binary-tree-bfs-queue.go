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

// BFS queue 
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0 
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	level := 0
	for len(queue) > 0 {
		size := len(queue)
		level++
		for i := 0; i < size; i++ {
			node := queue[i]
			// find first leave node
			if node.Left == nil && node.Right == nil {
				return level
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
	}
	return level
}
// @lc code=end

