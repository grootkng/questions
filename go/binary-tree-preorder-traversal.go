package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	arr := make([]int, 0)
	if root == nil {
		return arr
	}

	order(root, &arr)
	return arr
}

func order(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	*arr = append(*arr, root.Val)
	order(root.Left, arr)
	order(root.Right, arr)
}

func main() {
	farRoot := TreeNode{}
	farRoot.Val = 3
	farRoot.Left = nil
	farRoot.Right = nil

	midRoot := TreeNode{}
	midRoot.Val = 2
	midRoot.Left = &farRoot
	midRoot.Right = nil

	root := TreeNode{}
	root.Val = 1
	root.Left = nil
	root.Right = &midRoot

	fmt.Println(preorderTraversal(&root))
}
