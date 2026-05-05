package main

type ListNode struct {
    Val int
    Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil { return head; }

    // the stupid part of this question is that we dont know exactly how long the list is.

    length, last_node := linked_list_length_and_end_node(head);

    // if k > length, its more than one full rotation, dont care about extra
    k = k % length;

    // dont do anything.
    if k == 0 { return head; }

    new_tail := head;
    for range length - k - 1 { // advance until the node before the new node.
        new_tail = new_tail.Next;
    }

    new_head := new_tail.Next;
    if new_head == nil { panic("what"); }

    new_tail.Next = nil;
    last_node.Next = head;
    return new_head;
}

// returns the length, and the last node in the list
//
// last node == nil when head == nil
func linked_list_length_and_end_node(head *ListNode) (int, *ListNode) {
    result := 0;
    prev := head;
    for head != nil {
        result += 1;
        prev = head;
        head = head.Next;
    }
    return result, prev;
}
