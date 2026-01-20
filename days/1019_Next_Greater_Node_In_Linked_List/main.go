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
func nextLargerNodes(head *ListNode) []int {
	result := make([]int, 0)

	type Index_and_Num struct {
		num, index int
	}

	need_numbers := make([]Index_and_Num, 0)

	i := 0
	for head != nil {
		num := head.Val
		result = append(result, 0)

		for len(need_numbers) > 0 && need_numbers[len(need_numbers)-1].num < num {
			result[need_numbers[len(need_numbers)-1].index] = num
			need_numbers = need_numbers[:len(need_numbers)-1]
		}

		// add new number
		new_need_number := Index_and_Num{
			num:   num,
			index: i,
		}
		need_numbers = append(need_numbers, new_need_number)

		// advance.
		head = head.Next
		i += 1
	}

	return result
}
