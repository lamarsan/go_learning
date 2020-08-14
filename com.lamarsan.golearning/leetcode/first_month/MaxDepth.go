package first_month

/**
 * description 给定一个二叉树，找出其最大深度。
 * 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
 *
 * @author lamar
 * @date 2020/7/28 9:45 上午
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getMaxDepth(root, 0) + 1
}

func getMaxDepth(treeNode *TreeNode, depth int) int {
	leftDepth := depth
	rightDepth := depth
	if treeNode.Left != nil {
		leftDepth = getMaxDepth(treeNode.Left, depth+1)
	}
	if treeNode.Right != nil {
		rightDepth = getMaxDepth(treeNode.Right, depth+1)
	}
	return max(leftDepth, rightDepth)
}

func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
