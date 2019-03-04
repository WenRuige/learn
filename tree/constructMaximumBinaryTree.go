package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 最大二叉树
func constructMaximumBinaryTree(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	for i := 0; i < len(nums); i++ {

	}

	fmt.Println("hello world")

	return nil
}

func main() {
	constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
}
