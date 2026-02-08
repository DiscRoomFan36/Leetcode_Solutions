package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
    balanced, _ := is_balanced(root)
    return balanced
}

func is_balanced(node *TreeNode) (balanced bool, max_depth int) {
    if node == nil { return true, 0 }

    balanced_l, depth_l := is_balanced(node.Left)
    if !balanced_l { return false, 0 }
    balanced_r, depth_r := is_balanced(node.Right)
    if !balanced_r { return false, 0 }

    return Difference(depth_l, depth_r) <= 1, max(depth_l, depth_r)+1
}

func Difference(a, b int) int {
    if a <= b { return b - a
    } else    { return a - b }
}
