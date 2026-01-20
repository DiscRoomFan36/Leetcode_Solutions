package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	front := head
	back := head.Next

	for back != nil && back.Next != nil {
		node_to_add_to_front := back.Next
		node_to_add_to_back := back.Next.Next
		start_of_back_list := front.Next

		front.Next = node_to_add_to_front
		node_to_add_to_front.Next = start_of_back_list

		back.Next = node_to_add_to_back

		// ready for next iteration.
		front = node_to_add_to_front
		back = node_to_add_to_back
	}

	return head
}
