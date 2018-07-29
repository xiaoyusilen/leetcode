package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func isSameTree(p *TreeNode, q *TreeNode) bool {
	if (p == nil && q == nil) || (p != nil && q != nil && p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)) {
		return true
	}
	return false
}

func main() {
	fmt.Println(isSameTree(&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{1, nil, nil}}, &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{1, nil, nil}}))
}