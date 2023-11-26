package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if math.Abs(float64(height(root.Left, 0)-height(root.Right, 0))) <= 1 {
		return true && isBalanced(root.Left) && isBalanced(root.Right)
	} else {
		return false
	}
}

func height(root *TreeNode, level int) int {
	if root == nil {
		return level
	}
	level++
	sx := height(root.Left, level)
	dx := height(root.Right, level)
	if sx > dx {
		return sx
	}
	return dx
}
