package leetcode

import (
	"golang-acm/util"
)

/**
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
输入：head = [1,2,2,1]
输出：true

[1,2,3,2,1]


思路: 快慢指针，前面一半推入栈，后面一半比较前面出栈的数与后面遍历的数是否相等
 */

func IsPalindromeList2(head *util.ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return true
	}
	nums := []int{}
	temp := head
	for temp != nil {
		nums = append(nums, temp.Val)
		temp = temp.Next
	}
	n := len(nums)
	for i := 0; i < n/2; i++ {
		if nums[i] == nums[n-1-i] {
			continue
		} else {
			 return false
		}
	}
	return true
}

func IsPalindromeList(head *util.ListNode) bool {
	if head == nil {
		return false
	}

	prev := &util.ListNode{
		Next: head,
	}
	slow, fast := prev, prev

	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow

	numStack := []int{}
	prev = prev.Next
	for prev != mid.Next {
		numStack = append(numStack, prev.Val)
		prev = prev.Next
	}

	numStack2 := make([]int, len(numStack))
	copy(numStack2, numStack)
	numStack2 = append(numStack2, mid.Val)

	flag1, flag2 := true, true
	//mid = mid.Next
	for mid != nil {
		if len(numStack) >= 1 && mid.Val == numStack[len(numStack)-1] {
			mid = mid.Next
			if len(numStack) == 1 {
				numStack = []int{}
			} else {
				numStack = numStack[:len(numStack)-1]
			}
		} else {
			flag1 = false
			break
		}
	}

	mid2 := slow
	for mid2 != nil {
		if len(numStack2) >= 1 && mid2.Val == numStack2[len(numStack2)-1] {
			mid2 = mid2.Next
			if len(numStack2) == 1 {
				numStack2 = []int{}
			} else {
				numStack2 = numStack2[:len(numStack2)-1]
			}
		} else {
			flag2 = false
			break
		}
	}

	return flag1 || flag2
}