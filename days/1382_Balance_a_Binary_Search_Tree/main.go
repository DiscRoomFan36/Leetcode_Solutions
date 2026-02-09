package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
    if root == nil { return root; }

    sorted_array := make([]int, 0);
    binary_tree_to_sorted_array(root, &sorted_array);

    // now just sorted_array to binary tree
    return sorted_array_to_binary_tree(sorted_array);
}

func binary_tree_to_sorted_array(node *TreeNode, sorted_array *[]int) {
    if node == nil { return; }

    // everything left is smaller
    binary_tree_to_sorted_array(node.Left,  sorted_array);
    Append(sorted_array, node.Val);
    // everything right is bigger
    binary_tree_to_sorted_array(node.Right, sorted_array);
}

func sorted_array_to_binary_tree(sorted_array []int) *TreeNode {
    // the base case.
    if len(sorted_array) == 0 { return nil; }

    // are you ready for memory fragmentation?
    new_node := new(TreeNode);

    // split sorted_array into 2 pices,
    // turn sub pices into a sorted binary tree
    pivot_point := len(sorted_array) / 2;
    // middle value in array.
    new_node.Val = sorted_array[pivot_point];

    // could make this into a recursive function, but meh
    //
    // also quicksort moment
    new_node.Left  = sorted_array_to_binary_tree(sorted_array[              : pivot_point]);
    new_node.Right = sorted_array_to_binary_tree(sorted_array[pivot_point+1 :            ]);

    return new_node;
}

func Append[T any](slice *[]T, items ...T) {
	*slice = append(*slice, items...)
}