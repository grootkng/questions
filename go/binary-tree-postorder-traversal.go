package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	arr := make([]int, 0)
	order(root, &arr)
	return arr
}

func order(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	order(root.Left, arr)
	order(root.Right, arr)
	*arr = append(*arr, root.Val)
}

func main() {
	farRoot := TreeNode{}
	farRoot.Val = 2
	farRoot.Left = nil
	farRoot.Right = nil

	midRoot := TreeNode{}
	midRoot.Val = 1
	midRoot.Left = &farRoot
	midRoot.Right = nil

	root := TreeNode{}
	root.Val = 3
	root.Left = nil
	root.Right = &midRoot

	fmt.Println(postorderTraversal(&root))
}
