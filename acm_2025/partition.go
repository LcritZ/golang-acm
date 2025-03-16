package acm_2025

import "golang-acm/util"

func PartitionList(head *util.ListNode, x int) *util.ListNode {

	h1 := new(util.ListNode)
	h2 := new(util.ListNode)

	dummy1 := h1
	dummy2 := h2

	for head != nil {
		if head.Val < x {
			h1.Next = head
			h1 = h1.Next
		} else {
			h2.Next = h2.Next
			h2 = h2.Next
		}

		temp := head.Next
		head.Next = nil
		head = temp
	}
	if h1 != nil {
		h1.Next = dummy2.Next
		return dummy1.Next
	} else {
		return dummy2.Next
	}

}
