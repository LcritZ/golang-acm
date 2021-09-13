package basic

import "golang-acm/util"

/**
输入一颗二叉树的根节点和一个整数，按字典序打印出二叉树中结点值的和为输入整数的所有路径。
路径定义为从树的根结点开始往下一直到叶结点所经过的结点形成一条路径。
{10,5,12,4,7},22
[[10,5,7],[10,12]]

1 全局变量注意拷贝
2 深度遍历DFS赋值path
3 返回前需要path数组清空当前节点
4 添加节点root.val，sum－root.val, 再判断是否是叶子节点且sum==0
5 小于0则可以清空返回，大于可以继续向下追溯

 */

var path []int
var result [][]int

func FindPath( root *util.TreeNode,  expectNumber int ) [][]int {
	if root == nil {
		return result
	}
	DFSPath(root, expectNumber)
	return result
}

func DFSPath(root *util.TreeNode, sum int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	sum -= root.Val
	if 0 == sum && root.Left == nil && root.Right == nil {
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp)
		path = path[:len(path)-1]
		return
	}
	if sum > 0 {
		DFSPath(root.Left, sum)
		DFSPath(root.Right, sum)
	}

	path = path[:len(path)-1]
}
