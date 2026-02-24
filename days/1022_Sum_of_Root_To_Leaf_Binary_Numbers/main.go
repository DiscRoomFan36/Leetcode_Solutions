package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
    if root == nil { return 0; }
    return recur(root, 0);
}

// returns the sum
func recur(node *TreeNode, number int) int {
    new_number := (number << 1) + node.Val;

    if (node.Left == nil) && (node.Right == nil) {
        return new_number;
    }

    sum := 0;

    if node.Left  != nil { sum += recur(node.Left,  new_number); }
    if node.Right != nil { sum += recur(node.Right, new_number); }

    return sum;
}
