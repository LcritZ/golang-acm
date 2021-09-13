package leetcode

import "golang-acm/util"

/**
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

输入：head = [1], n = 1
输出：[]

输入：head = [1,2], n = 1
输出：[1]

 */

func RemoveNthFromEnd(head *util.ListNode, n int) *util.ListNode {
	if head == nil {
		return head
	}

	tmp := &util.ListNode{
		Next: head,
	}
	tail := tmp
	i := 1
	for i <= n {
		tail = tail.Next
		i++
	}
	if tail.Next == nil {
		return head.Next
	}
	for tail.Next != nil {
		tmp = tmp.Next
		tail = tail.Next
	}
	tmp.Next = tmp.Next.Next

	return head

}

func GF_removeNthFromEndStack(head *util.ListNode, n int) *util.ListNode {
	nodes := []*util.ListNode{}
	dummy := &util.ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next
}

func removeNthFromEnd(head *util.ListNode, n int) *util.ListNode {
	dummy := &util.ListNode{0, head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}
