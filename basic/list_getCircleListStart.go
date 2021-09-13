package basic

import "golang-acm/util"

func detectCycle(head *util.ListNode) *util.ListNode {
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

func detectCycle2(head *util.ListNode) *util.ListNode {
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
