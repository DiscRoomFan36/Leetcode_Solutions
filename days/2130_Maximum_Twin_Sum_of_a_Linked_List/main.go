package main

import "math"

type ListNode struct {
    Val int
    Next *ListNode
}

func pairSum(head *ListNode) int {
    length := List_Length(head);
    if length % 2 != 0 { panic("list is not even"); }

    sums := make([]int, length / 2);
    for i := range length / 2 {
        sums[i] = head.Val;
        head = head.Next;
    }

    for i := range length / 2 {
        sums[length/2 - 1 - i] += head.Val;
        head = head.Next;
    }

    best_sum := math.MinInt;
    for _, sum := range sums { best_sum = max(best_sum, sum); }
    return best_sum;
}

func List_Length(head *ListNode) int {
    count := 0;
    for head != nil {
        count += 1;
        head = head.Next;
    }
    return count;
}
