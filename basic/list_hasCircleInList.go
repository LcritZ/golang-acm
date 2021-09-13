package basic

import "golang-acm/util"

/**
链表是否有环
 */
func hasCircle(head *util.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
