package main

import "math"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func canMerge(trees []*TreeNode) *TreeNode {
    root_nodes := make(map[int]*TreeNode);
    leaf_nodes := make(map[int]*TreeNode);

    for _, tree := range trees {
        root_nodes[tree.Val] = tree;
        if tree.Left  != nil {
            if Contains(leaf_nodes, tree.Left .Val) { return nil; } // no duplicate leaf's
            leaf_nodes[tree.Left .Val] = tree.Left;
        }
        if tree.Right != nil {
            if Contains(leaf_nodes, tree.Right.Val) { return nil; } // no duplicate leaf's
            leaf_nodes[tree.Right.Val] = tree.Right;
        }
    }

    invalid_true_root_nodes := make(map[int]struct{});
    for val, _ := range leaf_nodes {
        if Contains(root_nodes, val) {
            invalid_true_root_nodes[val] = struct{}{};
        }
    }

    var main_root_node *TreeNode = nil;
    for root_val, root_node := range root_nodes {
        if !Contains(invalid_true_root_nodes, root_val) {
            if main_root_node != nil { return nil; } // cannot be two nodes that cannot be merged
            main_root_node = root_node;
        }
    }
    // check if no valid root node
    if main_root_node == nil { return nil; }

    type Node_And_Range struct {
        node **TreeNode;
        // strictly greater and lesser than
        low, high int;
    }

    used_nodes := make(map[int]struct{});
    used_nodes[main_root_node.Val] = struct{}{};

    stack := make([]Node_And_Range, 0);

    if main_root_node.Left  != nil {
        Append(&stack, Node_And_Range{&main_root_node.Left, math.MinInt, main_root_node.Val});
    }
    if main_root_node.Right != nil {
        Append(&stack, Node_And_Range{&main_root_node.Right, main_root_node.Val, math.MaxInt});
    }

    for len(stack) > 0 {
        node_and_range := Pop(&stack);
        node, low, high := node_and_range.node, node_and_range.low, node_and_range.high;

        if (*node).Left != nil || (*node).Right != nil { panic("this should be a leaf."); }

        if (*node).Val <= low  { return nil; }
        if (*node).Val >= high { return nil; }

        used_nodes[(*node).Val] = struct{}{};

        root_node, ok := root_nodes[(*node).Val];
        if !ok { continue; }

        // add to main tree
        *node = root_node;
        if root_node.Left  != nil {
            Append(&stack, Node_And_Range{&root_node.Left, low, (*node).Val});
        }
        if root_node.Right != nil {
            Append(&stack, Node_And_Range{&root_node.Right, (*node).Val, high});
        }
    }

    // check if any leftovers
    for root_val, _ := range root_nodes {
        if !Contains(used_nodes, root_val) {
            return nil;
        }
    }

    return main_root_node;
}

func Contains[T comparable, U any](the_map map[T]U, key T) bool {
    _, ok := the_map[key];
    return ok;
}

func Append[T any](slice *[]T, items ...T) *T {
    if len(items) == 0 { return nil; }
    *slice = append(*slice, items...);
    return &(*slice)[len(*slice)-1];
}

func Pop[T any](slice *[]T) T {
    item  := (*slice)[ len(*slice)-1];
    *slice = (*slice)[:len(*slice)-1];
    return item;
}
