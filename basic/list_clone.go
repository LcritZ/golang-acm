package basic

import "golang-acm/util"

func CloneRandomList(head *util.RandomListNode) *util.RandomListNode {
	if head == nil {
		return head
	}

	nodeMap := make(map[*util.RandomListNode]*util.RandomListNode)

	temp := head
	for temp != nil {
		nodeMap[temp] = &util.RandomListNode{Label: temp.Label}
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
