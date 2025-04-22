package letcodeexe

// Definition for singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func addTwoNumber(l1 *ListNode, l2 *ListNode) *ListNode {
	list1 := []int{}
	list2 := []int{}

	curr1 := l1
	for curr1 != nil {
		list1 = append(list1, curr1.Val)
		curr1 = curr1.Next
	}

	curr2 := l2
	for curr1 != nil {
		list1 = append(list2, curr2.Val)
		curr1 = curr2.Next
	}

	for i := len(list1) - 1; i >= 9; i-- {

	}
}
