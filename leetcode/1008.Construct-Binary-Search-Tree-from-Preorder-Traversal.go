package leetcode

import (
	"math"

	. "github.com/AKA-bingo/GoCoPark/base"
)

/**
Return the root node of a binary search tree that matches the given preorder traversal.

(Recall that a binary search tree is a binary tree where for every node, any descendant of node.left has a value < node.val, and any descendant of node.right has a value > node.val.  Also recall that a preorder traversal displays the value of the node first, then traverses node.left, then traverses node.right.)

Example 1:
Input: [8,5,1,7,10,12]
Output: [8,5,10,1,7,null,12]

Note:

1 <= preorder.length <= 100
The values of preorder are distinct.
*/

func BstFromPreorderBak(preorder []int) *TreeNode {
	if len(preorder) <= 0 {
		return nil
	}

	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	var leftPos int
	for i, num := range preorder[1:] {
		if num > preorder[0] {
			leftPos = i
			break
		}
	}

	root.Left = BstFromPreorderBak(preorder[1:leftPos])
	root.Right = BstFromPreorderBak(preorder[leftPos:])
	return root
}

func BstFromPreorder(preorder []int) *TreeNode {
	return bstBuild(preorder, new(int), math.MaxInt32)
}

func bstBuild(preorder []int, index *int, rootValue int) *TreeNode {
	if *index >= len(preorder) || preorder[*index] > rootValue {
		return nil
	}
	node := &TreeNode{
		Val: preorder[*index],
	}
	*index++

	node.Left = bstBuild(preorder, index, node.Val)
	node.Right = bstBuild(preorder, index, rootValue)
	return node
}
