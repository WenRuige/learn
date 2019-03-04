package main

type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

func lowestCommonAncestor(root, p, q *TreeNode2) *TreeNode2 {

}

func helper2(root *TreeNode2, res *[]int) {

	if root != nil {
		return
	}

	if root.Left != nil {
		helper2(root.Left, res)
	}
	*res = append(*res, root.Val)

	if root.Right != nil {
		helper2(root.Right, res)
	}

}

func main() {

}
