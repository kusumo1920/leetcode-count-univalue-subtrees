package main

import "fmt"

func main() {
	input := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	output := countUnivalSubtrees(input)
	fmt.Println(output)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countUnivalSubtrees(root *TreeNode) int {
	count := 0
	var recursiveFn func(node *TreeNode) bool
	recursiveFn = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		left := recursiveFn(node.Left)
		right := recursiveFn(node.Right)

		if left && right {
			if node.Left != nil && node.Left.Val != node.Val {
				return false
			}
			if node.Right != nil && node.Right.Val != node.Val {
				return false
			}
			count++
			return true
		}
		return false
	}

	recursiveFn(root)

	return count
}
