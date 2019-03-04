package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {

	result := []int{}
	helper(root, &result)
	return result[k-1]
}

// 遍历顺序,左根右
func helper(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		helper(root.Left, res)
	}
	*res = append(*res, root.Val)
	if root.Right != nil {
		helper(root.Right, res)
	}
}

func main() {
	root := new(TreeNode)
	root.Val = 3
	root.Left = new(TreeNode)
	root.Left.Val = 1
	root.Right = new(TreeNode)
	root.Right.Val = 4
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 2

	res := kthSmallest(root, 1)
	fmt.Println(res)

}
