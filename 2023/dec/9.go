package dec

// Problem: https://leetcode.com/problems/binary-tree-inorder-traversal/description
/*
Input: root = [1,null,2,3]
Output: [1,3,2]
*/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	return inorder(root, nil)
  }
  
  func inorder(root *TreeNode, res []int) []int {
	if root == nil {
	    return res
	}
  
	res = inorder(root.Left, res)
	res = append(res, root.Val)
	res = inorder(root.Right, res)
	
	return res
  }