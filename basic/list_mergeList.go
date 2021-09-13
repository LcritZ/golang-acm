package basic

import "golang-acm/util"

func Merge( pHead1 *util.ListNode,  pHead2 *util.ListNode) *util.ListNode {
	// write code here
	if pHead1 == nil && pHead2 == nil {
		return nil
	} else if pHead1 == nil {
		return pHead2
	} else if pHead2 == nil {
		return pHead1
	}

	if pHead1.Val <= pHead2.Val {
		pHead1.Next= Merge(pHead1.Next,pHead2);
		return pHead1;
	} else if pHead1.Val == pHead2.Val {
		temp := pHead1.Next
		pHead1.Next = pHead2
		pHead1.Next.Next = Merge(temp,pHead2.Next);
		return pHead1;
	} else {
		pHead2.Next= Merge(pHead1,pHead2.Next);
		return pHead2;
	}
}


func Merge2( pHead1 *util.ListNode,  pHead2 *util.ListNode) *util.ListNode {
	if pHead1 == nil && pHead2 == nil {
		return nil
	} else if pHead1 == nil {
		return pHead2
	} else if pHead2 == nil {
		return pHead1
	}
	h := &util.ListNode{}
	temp := h
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val <= pHead2.Val {
			temp.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			temp.Next = pHead2
			pHead2 = pHead2.Next
		}
		temp = temp.Next
	}
	if pHead1 == nil {
		temp.Next = pHead2
	}
	if pHead2 == nil {
		temp.Next = pHead1
	}
	return h.Next

}

