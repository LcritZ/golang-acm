package basic

import "golang-acm/util"

func DeleteNode(pHead *util.ListNode, val int) *util.ListNode {
	if pHead == nil {
		return pHead
	}
	var temp *util.ListNode = nil
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
