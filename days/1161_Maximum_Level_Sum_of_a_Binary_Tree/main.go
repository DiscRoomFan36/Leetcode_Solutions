package main

import "math"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
    if root == nil { panic("no nil roots allowed"); }

    curr_stack := []*TreeNode{root};
    next_stack := []*TreeNode{};

    best_level_sum := math.MinInt;
    best_level := 1

    current_level := 1;
    for len(curr_stack) != 0 {

        level_sum := 0;
        for len(curr_stack) > 0 {
            node := Pop(&curr_stack);

            level_sum += node.Val;

            if node.Left  != nil { Append(&next_stack, node.Left ); }
            if node.Right != nil { Append(&next_stack, node.Right); }
        }

        if level_sum > best_level_sum {
            best_level_sum = level_sum;
            best_level = current_level;
        }

        // swap the stacks.
        curr_stack, next_stack = next_stack, curr_stack;
        current_level += 1;
    }

    return best_level;
}

func Pop[T any](slice *[]T) T {
    item  := (*slice)[ len(*slice)-1]
    *slice = (*slice)[:len(*slice)-1]
    return item
}

func Append[T any](slice *[]T, items ...T) {
	*slice = append(*slice, items...)
}
