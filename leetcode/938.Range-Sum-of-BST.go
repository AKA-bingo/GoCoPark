package leetcode

import (
	. "github.com/AKA-bingo/GoCoPark/base"
)

/*
Given the root node of a binary search tree, return the sum of values of all nodes with value between L and R (inclusive).

The binary search tree is guaranteed to have unique values.

Example 1:
Input: root = [10,5,15,3,7,null,18], L = 7, R = 15
Output: 32

Example 2:
Input: root = [10,5,15,3,7,13,18,1,null,6], L = 6, R = 10
Output: 23

Note:
The number of nodes in the tree is at most 10000.
The final answer is guaranteed to be less than 2^31.
*/

func RangeSumBST(root *TreeNode, L int, R int) int {
	var sum int
	var countFunc func(*TreeNode)

	countFunc = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Val >= L && node.Val <= R {
			sum += node.Val
		}
		if node.Val <= L {
			countFunc(node.Right)
		} else if node.Val >= R {
			countFunc(node.Left)
		} else {
			countFunc(node.Right)
			countFunc(node.Left)
		}
	}

	countFunc(root)
	return sum
}
