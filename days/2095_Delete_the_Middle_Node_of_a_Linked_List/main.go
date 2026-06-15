package main

type ListNode struct {
    Val int
    Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
    // what your suppost to do is 2 pointers method of obtaining the middle of a linked list,
    // I however, do not care.
    //
    // im just gonna measure the length of the array.
    length := List_Length(head);
    if length <= 1 { return nil; }

    node := head;
    for range length / 2 - 1 { node = node.Next; }

    node.Next = node.Next.Next;
    return head;
}

func List_Length(head *ListNode) int {
    count := 0;
    for head != nil {
        count += 1;
        head = head.Next;
    }
    return count;
}
