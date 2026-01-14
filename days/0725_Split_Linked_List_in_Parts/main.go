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
func splitListToParts(head *ListNode, k int) []*ListNode {
	result := make([]*ListNode, k)

	length := 0
	for p := head; p != nil; p = p.Next {
		length += 1
	}

	for i := range k {
		num_to_take := length / k
		if i < length%k {
			num_to_take += 1
		}

		if num_to_take > 0 {
			result[i] = head

			for range num_to_take - 1 {
				head = head.Next
			}

			new_head := head.Next
			head.Next = nil
			head = new_head
		}
	}

	return result
}
