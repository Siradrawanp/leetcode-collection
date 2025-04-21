package letcodeexe

import "sort"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	vals1 := []int{}
	vals2 := []int{}

	curr1 := list1
	for curr1 != nil {
		vals1 = append(vals1, curr1.Val)
		curr1 = curr1.Next
	}

	curr2 := list2
	for curr2 != nil {
		vals2 = append(vals2, curr2.Val)
		curr2 = curr2.Next
	}

	vals1 = append(vals1, vals2...)
	sort.Ints(vals1)

	var head, tail *ListNode
	for _, num := range vals1 {
		newNode := &ListNode{
			Val:  num,
			Next: nil,
		}

		if head == nil {
			head = newNode
			tail = newNode
		} else {
			tail.Next = newNode
			tail = newNode
		}
	}

	return head
}
