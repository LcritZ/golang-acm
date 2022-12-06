package leetcode

import (
	"golang-acm/util"
)

/**
给定一个头结点为 head 的非空单链表，返回链表的中间结点。

如果有两个中间结点，则返回第二个中间结点。
2,3,4,6,7,9,10

2,3,4,6,7,9
2
2,3
 */
func middleNode(head *util.ListNode) *util.ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func middleNode2(head *util.ListNode) *util.ListNode {
	dummy := &util.ListNode{Next: head}
	slow, fast := dummy, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow.Next
}
