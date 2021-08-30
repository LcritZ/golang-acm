package leetcode

import (
	"golang-acm/common"
)

func GF_AddTwoNumbers(l1, l2 *common.ListNode) (head *common.ListNode) {
	var tail *common.ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &common.ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &common.ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &common.ListNode{Val: carry}
	}
	return
}

func AddTwoNumbers(l1, l2 *common.ListNode) *common.ListNode {
	var tail *common.ListNode
	var head *common.ListNode
	carry := 0

	for l1 != nil || l2 != nil {
		v1, v2, sum := 0, 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum = (v1+v2+carry)%10
		carry = (v1+v2+carry)/10
		if head == nil {
			head = &common.ListNode{
				Val: sum,
			}
			tail = head
		} else {
			tail.Next = &common.ListNode{
				Val: sum,
			}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &common.ListNode{
			Val: carry,
		}
	}

	return head
}

func AddTwoNumbers2(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	var head, tail *common.ListNode
	carry := 0

	for l1 != nil || l2 != nil {
		v1, v2, sum := 0, 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum = (v1 + v2 + carry)%10
		carry = (v1 + v2 + carry)/10
		if head == nil {
			 head = &common.ListNode{
			 	Val: sum,
			 }
			 tail = head
		} else {
			tail.Next = &common.ListNode{
				Val: sum,
			}
			tail = tail.Next
		}
	}

	if carry > 0 {
		tail.Next = &common.ListNode{
			Val: carry,
		}
	}

	return head
}


