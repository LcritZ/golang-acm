package basic

import "golang-acm/util"

func FindKthToTail( pHead *util.ListNode,  k int ) *util.ListNode {
	// write code here
	if k <= 0 {
		return nil
	}
	if pHead == nil {
		return nil
	}

	var tailNode *util.ListNode = pHead
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