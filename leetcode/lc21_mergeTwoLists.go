package leetcode

import (
	"golang-acm/util"
)

/**
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

*/

func MergeTwoLists(l1 *util.ListNode, l2 *util.ListNode) *util.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val <= l2.Val {
		l1.Next = MergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = MergeTwoLists(l1, l2.Next)
		return l2
	}

}

func MergeTwoLists2(l1 *util.ListNode, l2 *util.ListNode) *util.ListNode {
	dummyNode := &util.ListNode{}
	temp := dummyNode
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			temp.Next = l1
			l1 = l1.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
		}
		temp = temp.Next
	}

	if l1 == nil {
		temp.Next = l2
	} else {
		temp.Next = l1
	}

	return dummyNode.Next

}

func MergeTwoLists3(l1 *util.ListNode, l2 *util.ListNode) *util.ListNode {
	dummy := &util.ListNode{}
	temp := dummy

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			temp.Next = l1
			l1 = l1.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
		}
		temp = temp.Next
	}

	if l1 == nil {
		temp.Next = l2
	} else {
		temp.Next = l1
	}

	return dummy.Next

}