package main

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil { return list2; }
    if list2 == nil { return list1; }

    // both lists have at least 1 element.

    var head, tail *ListNode;
    if list1.Val <= list2.Val {
        head = list1;
        tail = list1;
        list1 = list1.Next
    } else {
        head = list2;
        tail = list2;
        list2 = list2.Next;
    }

    for {
        if list1 == nil {
            tail.Next = list2;
            break;
        }
        if list2 == nil {
            tail.Next = list1;
            break;
        }

        if list1.Val <= list2.Val {
            tail.Next = list1;
            list1 = list1.Next;
        } else {
            tail.Next = list2;
            list2 = list2.Next;
        }
        tail = tail.Next;
    }

    return head;
}
