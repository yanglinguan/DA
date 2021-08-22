/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
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
func buildTree(preorder []int, inorder []int) *TreeNode {
    inorderIdx := map[int]int{}
		for i, v := range inorder {
			inorderIdx[v] = i
		}
		return helper(preorder, inorder, 0, 0, len(inorder)-1, inorderIdx)
}

func helper(preorder, inorder []int, preIdx, il, ir int, inorderIdx map
	[int]int) *TreeNode {
	if il > ir {
		return nil
	}

	root := &TreeNode{preorder[preIdx], nil, nil}
	rootIdx := inorderIdx[preorder[preIdx]]
	// find the number of nodes in the left tree
	leftNum := rootIdx - il
	root.Left = helper(preorder, inorder, preIdx+1, il, rootIdx-1, inorderIdx)
	// right tree starts at preId+1+leftNum
	root.Right = helper(preorder, inorder, preIdx+1+leftNum, rootIdx+1, ir, inorderIdx)
	return root
}
// @lc code=end

