package acm

import "golang-acm/common"

func CloneRandomList(head *common.RandomListNode) *common.RandomListNode {
	if head == nil {
		return head
	}

	nodeMap := make(map[*common.RandomListNode]*common.RandomListNode)

	temp := head
	for temp != nil {
		nodeMap[temp] = &common.RandomListNode{Label: temp.Label}
		temp = temp.Next
	}

	newHead := nodeMap[head]

	temp = head
	for temp != nil {
		if temp.Next != nil {
			newHead.Next = nodeMap[temp.Next]
		}
		if temp.Random != nil {
			newHead.Random = nodeMap[temp.Random]
		}
		temp = temp.Next
		newHead = newHead.Next
	}

	return nodeMap[head]
}
