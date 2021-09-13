package util

type ListNode struct{
	Val int
	Next *ListNode
}

type DobuleListNode struct {
	Val int
	Next *DobuleListNode
	Pre  *DobuleListNode
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type RandomListNode struct {
	Label int
	Next *RandomListNode
	Random *RandomListNode
}

func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}