package acm

import (
	"fmt"
	"golang-acm/common"
)

var ch = make(chan int, 3)

func Produce(input int) {
	ch <-input
	fmt.Println("produce:", input)
}

func Consumer() int {
	for  {
		select {
		case output :=<-ch:
			fmt.Println("consumer:", output)
			return output
		default:
			fmt.Printf("no left")
			return -1
		}
	}

}

var totalPath [][]int
var subpath []int

/**
root节点到叶子  sum=k 的节点path
 */
func FindPathK(root *common.TreeNode, k int) [][]int {
	if root == nil {
		return totalPath
	}
	DFSPathK(root, k)
	return totalPath
}

func DFSPathK(node *common.TreeNode, k int) {

	if k-node.Val == 0 {
		if node.Right == nil && node.Left == nil  {
			subpath = append(subpath, node.Val)
			totalPath = append(totalPath, subpath)
		}
	}

	if k-node.Val < 0 {
		return
	} else {
		subpath = append(subpath, node.Val)
		if node.Left != nil {
			DFSPathK(node.Left, k-node.Val)
		}
		if node.Right != nil {
			DFSPathK(node.Left, k-node.Val)
		}
	}
	//if len(subpath) > 1 {
	//
	//}
	subpath = subpath[:len(subpath)-1]

}



