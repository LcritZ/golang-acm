package acm

import "common"

func DeleteNode(pHead *common.ListNode, val int) *common.ListNode {
	if pHead == nil {
		return pHead
	}
	var temp *common.ListNode = nil
	if pHead.Val == val {
		temp = pHead.Next
	} else {
		temp = pHead
		for pHead.Next != nil{
			if pHead.Next.Val == val {
				pHead.Next = pHead.Next.Next
				break
			}
			pHead = pHead.Next
		}
	}

	return temp
}
