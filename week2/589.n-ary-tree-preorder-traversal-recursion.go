/*
 * @lc app=leetcode id=589 lang=golang
 *
 * [589] N-ary Tree Preorder Traversal
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// 1. recursion 

func preorder(root *Node) []int {
	// base case
	if root == nil {
		return []int{}
	}

	res := []int{root.Val}
	for _, c := range root.Children {
		res = append(res, preorder(c)...)
	}
	return res
}
// @lc code=end

