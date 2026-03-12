package main

type ListNode struct {
    Val int
    Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
    nums_as_a_set := make(map[int]struct{});
    for _, n := range nums { nums_as_a_set[n] = struct{}{}; }

    // remove until we find a clean node
    new_head := head;
    for new_head != nil && Contains(nums_as_a_set, new_head.Val) {
        new_head = new_head.Next;
    }
    if new_head == nil { panic("this is a leetcode question, if this were real i would return nil."); }

    tail := new_head;
    for tail.Next != nil {
        // check if this next thing is a bad node.
        if Contains(nums_as_a_set, tail.Next.Val) {
            // was a bad node, skip it.
            tail.Next = tail.Next.Next;
        } else {
            // was a good node, keep it, move tail.
            tail = tail.Next;
        }
    }

    return new_head;
}

func Contains[T comparable, U any](the_map map[T]U, item T) bool {
    _, found := the_map[item];
    return found;
}