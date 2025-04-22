package letcodeexe

// Definition for singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func addTwoNumber(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode

	sisa := 0
	for l1 != nil && l2 != nil {
		num := (l1.Val + l2.Val) + sisa
		if num > 9 {
			sisa = 1
		} else {
			sisa = 0
		}

		// add to linked list
		ins := num % 10
		newNode := &ListNode{
			Val:  ins,
			Next: nil,
		}

		if head == nil {
			head = newNode
			tail = newNode
		} else {
			tail.Next = newNode
			tail = newNode
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		num := l1.Val + sisa
		if num > 9 {
			sisa = 1
		} else {
			sisa = 0
		}

		// add to linked list
		ins := num % 10
		newNode := &ListNode{
			Val:  ins,
			Next: nil,
		}

		if head == nil {
			head = newNode
			tail = newNode
		} else {
			tail.Next = newNode
			tail = newNode
		}
		l1 = l1.Next
	}

	for l2 != nil {
		num := l2.Val + sisa
		if num > 9 {
			sisa = 1
		} else {
			sisa = 0
		}

		// add to linked list
		ins := num % 10
		newNode := &ListNode{
			Val:  ins,
			Next: nil,
		}

		if head == nil {
			head = newNode
			tail = newNode
		} else {
			tail.Next = newNode
			tail = newNode
		}
		l2 = l2.Next
	}

	if sisa != 0 {
		// add to linked list
		newNode := &ListNode{
			Val:  sisa,
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
