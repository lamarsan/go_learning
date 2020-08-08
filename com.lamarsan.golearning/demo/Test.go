package main

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(f(root), g(root))
}

func f(o *TreeNode) int {
	if o == nil {
		return 0
	}
	return g(o.Left) + g(o.Right) + o.Val
}

func g(o *TreeNode) int {
	if o == nil {
		return 0
	}
	return max(f(o.Left), g(o.Left)) + max(f(o.Right), g(o.Right))
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a := "123"
	b := "321"
	c := a + b
	println(c)
}
