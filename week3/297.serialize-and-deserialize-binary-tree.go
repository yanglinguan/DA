/*
 * @lc app=leetcode id=297 lang=golang
 *
 * [297] Serialize and Deserialize Binary Tree
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

type Codec struct {
    
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "nil"
	}

	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)

}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
	dataList := strings.Split(data, ",")
	idx := 0
	return helper(dataList, &idx)
}

func helper(data []string, idx *int) *TreeNode {
	if data[*idx] == "nil" {
		return nil
	}
	v, _ := strconv.Atoi(data[*idx])
	root := &TreeNode{v, nil, nil}
	*idx++
	root.Left = helper(data, idx)
	*idx++
	root.Right = helper(data, idx)
	return root
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end

