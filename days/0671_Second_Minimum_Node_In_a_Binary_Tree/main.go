package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
    if root == nil { return -1; }

    ok, num := smallest_bigger_than(root, root.Val);
    if !ok { return -1; }
    return num;
}

func smallest_bigger_than(root *TreeNode, num int) (bool, int) {
    if root == nil    { return false, 0;  }
    if root.Val > num { return true, root.Val; }

    ok1, num1 := smallest_bigger_than(root.Left,  num);
    ok2, num2 := smallest_bigger_than(root.Right, num);

    if !ok1 && !ok2 { return false, 0; }
    if ok1 && !ok2 {
        return true, num1;
    }
    if !ok1 && ok2 {
        return true, num2;
    }
    return true, min(num1, num2);
}
