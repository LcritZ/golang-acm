package acm_2025

import (
	"golang-acm/util"
)

func MergeTwoLists(list1 *util.ListNode, list2 *util.ListNode) *util.ListNode {

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	head := new(util.ListNode)
	dummy := head

	h1 := list1
	h2 := list2

	for h1 != nil && h2 != nil {
		if h1.Val <= h2.Val {
			dummy.Next = h1
			h1 = h1.Next
		} else {
			dummy.Next = h2
			h2 = h2.Next
		}
		dummy = dummy.Next
	}

	if h1 == nil {
		dummy.Next = h2
	} else if h2 == nil {
		dummy.Next = h1
	}

	return head.Next
}
