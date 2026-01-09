
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    if root == nil { return []int{} }

    result := make([]int, 0)

    type Stack_Item struct {
        seen_before bool
        node *TreeNode
    }

    stack := make([]Stack_Item, 1)
    stack[0] = Stack_Item{false, root}

    for len(stack) > 0 {
        // pop
        item := &stack[len(stack)-1]
        if !item.seen_before {
            item.seen_before = true

            if item.node.Right != nil { stack = append(stack, Stack_Item{false, item.node.Right}) }
            if item.node.Left  != nil { stack = append(stack, Stack_Item{false, item.node.Left }) }

        } else {
            result = append(result, item.node.Val)
            stack = stack[:len(stack)-1]
        }
    }

    return result
}
