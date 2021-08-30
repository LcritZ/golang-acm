package leetcode

import (
	"golang-acm/common"
)

/**
2,3
2,3,4,6
链表重排
 */

func ReorderList(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil  {
		return head
	}
	mid := getMidNode(head)

	h1 := head
	h2 := reverseList(mid.Next)

	mid.Next = nil

	mergeList(h1, h2)
	return h1
}

func getMidNode(head *common.ListNode) *common.ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *common.ListNode) *common.ListNode {
	var tail *common.ListNode = nil
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = tail
		tail = curr
		curr = next
	}
	return tail
}

func mergeList(h1, h2 *common.ListNode) *common.ListNode {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}

	temp := h1.Next
	h1.Next = h2
	h1.Next.Next = mergeList(temp, h2.Next)
	return h1
}

/**
使用数组保存。然后拼接
2,3，4
2,3,4,6
 */
func reorderList2(head *common.ListNode) {
	if head == nil {
		return
	}
	nodes := []*common.ListNode{}
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}
