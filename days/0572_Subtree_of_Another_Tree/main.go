package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if is_equal(root, subRoot) { return true; }

    if root == nil { return false; }

    if root.Left != nil {
        if isSubtree(root.Left, subRoot) { return true; }
    }

    if root.Right != nil {
        if isSubtree(root.Right, subRoot) { return true; }
    }

    return false;
}

func is_equal(a *TreeNode, b *TreeNode) bool {
    if a == b { return true; }

    if a == nil { return false; }
    if b == nil { return false; }

    if a.Val != b.Val { return false; }

    if !is_equal(a.Left, b.Left)   { return false; }
    if !is_equal(a.Right, b.Right) { return false; }
    return true;
}
