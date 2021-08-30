package acm

import (
	"fmt"
	"golang-acm/common"
)

func nowcodeReverseList( pHead *common.ListNode) *common.ListNode {
	// write code here
	if pHead == nil {
		return pHead
	}

	var temp *common.ListNode
	var newHead *common.ListNode = nil
	for pHead != nil {
		temp = pHead.Next
		pHead.Next = newHead
		newHead = pHead
		pHead = temp
	}

	return newHead

}

// 1 2 3 4
// 4 3 2 1
func ReverseDoubleList(pHead *common.DobuleListNode) *common.DobuleListNode {
	if pHead.Next == nil {
		return pHead
	}
	var temp *common.DobuleListNode
	var newHead *common.DobuleListNode = nil
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

func main() {
	var pHead *common.DobuleListNode = &common.DobuleListNode{
		Val: 1,
	}
	var p2 *common.DobuleListNode = &common.DobuleListNode{
		Val: 2,
	}
	var p3 *common.DobuleListNode = &common.DobuleListNode{
		Val: 3,
	}
	pHead.Next = p2
	p2.Pre = pHead
	p2.Next = p3
	p3.Pre = p2
	fmt.Println(ReverseDoubleList(pHead))
}