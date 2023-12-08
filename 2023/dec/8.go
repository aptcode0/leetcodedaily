package dec

import "fmt"

// Problem: https://leetcode.com/problems/construct-string-from-binary-tree/description
/*
Input: root = [1,2,3,4]
Output: "1(2(4))(3)"
*/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	if root == nil {
	    return ""
	}
	leftStr := tree2str(root.Left)
	rightStr := tree2str(root.Right)
  
	ans := fmt.Sprint(root.Val)
	if len(rightStr) > 0 {
	    ans = fmt.Sprintf("%s(%s)(%s)", ans, leftStr, rightStr)
	} else if len(leftStr) > 0 {
	    ans = fmt.Sprintf("%s(%s)", ans, leftStr)    
	}
  
	return ans
  }