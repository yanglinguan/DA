/*
 * @lc app=leetcode id=429 lang=golang
 *
 * [429] N-ary Tree Level Order Traversal
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
// BFS
func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)
	res := make([][]int, 0)
	for len(queue) > 0 {
		// record the size, only loop the node in this level
		size := len(queue)
		// record the values in this level
		level := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[i]
			level[i] = node.Val
			for _, c := range node.Children {
				queue = append(queue, c)
			}
		}
		queue = queue[size:]
		res = append(res, level)
	}
	return res
}
// @lc code=end

