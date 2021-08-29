/*
 * @lc app=leetcode id=515 lang=golang
 *
 * [515] Find Largest Value in Each Tree Row
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
func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	res := []int{}
	for len(queue) > 0 {
		size := len(queue) 
		max := math.MinInt32
		for i := 0; i < size; i++ {
			n := queue[i]
			if n.Val > max {
				max = n.Val
			}
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
		queue = queue[size:]
		res = append(res, max)
	}
	return res
}
// @lc code=end

