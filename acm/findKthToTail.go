package acm

import "golang-acm/common"

func FindKthToTail( pHead *common.ListNode,  k int ) *common.ListNode {
	// write code here
	if k <= 0 {
		return nil
	}
	if pHead == nil {
		return nil
	}

	var tailNode *common.ListNode = pHead
	for i := 0; i < k-1; i++ {
		if tailNode.Next == nil {
			return nil
		} else {
			tailNode = tailNode.Next
		}
	}
	for tailNode.Next != nil {
		pHead = pHead.Next
		tailNode = tailNode.Next
	}

	return pHead
}