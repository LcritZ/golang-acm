package acm

import "golang-acm/common"

func detectCycle(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			temp := head
			for temp != slow {
				temp = temp.Next
				slow = slow.Next
			}
			return temp
		}
	}
	return nil
}

func detectCycle2(head *common.ListNode) *common.ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
