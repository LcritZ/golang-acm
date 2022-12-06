package basic

import (
	"golang-acm/util"
)

func nowcodeReverseList( pHead *util.ListNode) *util.ListNode {
	// write code here
	if pHead == nil {
		return pHead
	}

	//var temp *util.ListNode
	var newHead *util.ListNode = nil
	for pHead != nil {
		temp := pHead.Next
		pHead.Next = newHead
		newHead = pHead
		pHead = temp
	}

	return newHead
}

// 1 2 3 4
// 4 3 2 1
func ReverseDoubleList(pHead *util.DobuleListNode) *util.DobuleListNode {
	if pHead.Next == nil {
		return pHead
	}
	var temp *util.DobuleListNode
	var newHead *util.DobuleListNode = nil
	for pHead != nil {
		temp = pHead.Next
		pHead.Pre = temp
		pHead.Next = newHead
		newHead = pHead
		pHead = temp
	}
	newHead.Pre = nil
	return newHead
}
