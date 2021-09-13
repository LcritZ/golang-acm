package leetcode

import (
	"golang-acm/util"
)

func ReverseList(head *util.ListNode) *util.ListNode {
	if head == nil {
		return head
	}

	var prev *util.ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev

}

func ReverseList2(head *util.ListNode) *util.ListNode {
	if head.Next == nil {
		return head
	}

	last := ReverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}
