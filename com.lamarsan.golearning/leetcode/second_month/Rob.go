package second_month

/**
 * description 在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。
 * 这个地区只有一个入口，我们称之为“根”。 除了“根”之外，每栋房子有且只有一个“父“房子与之相连。
 * 一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。
 * 如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。
 * 计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。
 *
 * @author lamar
 * @date 2020/8/5 10:01 上午
 */
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

// ------------------------------------------------------------------------------
func rob2(root *TreeNode) int {
	val := dfs(root)
	return max(val[0], val[1])
}

func dfs(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	l, r := dfs(node.Left), dfs(node.Right)
	selected := node.Val + l[1] + r[1]
	notSelected := max(l[0], l[1]) + max(r[0], r[1])
	return []int{selected, notSelected}
}
