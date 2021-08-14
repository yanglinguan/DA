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
// iteration

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	stack := make([]colorNode, 0)
	stack = append(stack, colorNode{0, root})
	res := make([]int, 0)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		if top.color == 0 {
			top.color = 1
			// add the stack in the reverse order
			// since stack is LIFO
			for i := len(top.node.Children) - 1; i >= 0; i-- {
				stack = append(stack, colorNode{0, top.node.Children[i]})
			}
			stack = append(stack, top)
		} else {
			res = append(res, top.node.Val)
		}
	}
	return res
}

type colorNode struct {
	// 0: not visited
	// 1: visited, can add to result set
	color int
	node *Node
}
// @lc code=end

