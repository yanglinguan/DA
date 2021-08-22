/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	inorderMap := make(map[int]int)
	for i, n := range inorder {
		inorderMap[n] = i
	}

	return helper(inorder, postorder, len(postorder)-1, 0, len(inorder)-1, inorderMap)

}

func helper(inorder, postorder []int, pIdx, il, ir int, iMap map[int]int) *TreeNode {
	if il > ir {
		return nil
	}

	root := &TreeNode{postorder[pIdx], nil, nil}
	rootIdx := iMap[postorder[pIdx]]
	rightNum := ir - rootIdx
	root.Left = helper(inorder, postorder, pIdx-1-rightNum, il, rootIdx-1, iMap)
	root.Right = helper(inorder, postorder, pIdx-1, rootIdx+1, ir, iMap)
	return root
}
// @lc code=end

