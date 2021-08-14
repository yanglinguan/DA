/*
 * @lc app=leetcode id=590 lang=golang
 *
 * [590] N-ary Tree Postorder Traversal
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	for _, c := range root.Children {
		res = append(res, postorder(c)...)
	}
	res = append(res, root.Val)
	return res
}
// @lc code=end

