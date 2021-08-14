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
	stack := []colorNode{}
	res := []int{}
	if root == nil { return res }
	
	stack = append(stack, colorNode{0, root})
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top.color == 0 {
			top.color = 1
			stack = append(stack, top)
			for i := len(top.node.Children)-1; i >= 0; i-- {
				stack = append(stack, colorNode{0, top.node.Children[i]})
			}
		} else {
			res = append(res, top.node.Val)
		}
	}
	return res
}

type colorNode struct {
	color int
	node *Node
}
// @lc code=end

