package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    if root == nil { return []int{} }

    result := make([]int, 0)

    type Item struct {
        node *TreeNode
        seen_before bool
    }

    stack := make([]Item, 0)
    Append(&stack, Item{root, false})

    for len(stack) > 0 {
        item := Pop(&stack)
        if !item.seen_before {
            Append(&stack, Item{item.node, true})
            if item.node.Left != nil {
                Append(&stack, Item{item.node.Left, false})
            }
        } else {
            Append(&result, item.node.Val)
            if item.node.Right != nil {
                Append(&stack, Item{item.node.Right, false})
            }
        }
    }

    return result
}

func Append[T any](slice *[]T, items ...T) {
	*slice = append(*slice, items...)
}

func Pop[T any](slice *[]T) T {
	item  := (*slice)[ len(*slice)-1]
	*slice = (*slice)[:len(*slice)-1]
	return item
}
