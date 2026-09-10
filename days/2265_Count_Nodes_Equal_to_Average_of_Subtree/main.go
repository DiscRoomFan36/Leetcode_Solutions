package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
    type Sum_And_Count struct {
        sum int;
        count int;
    }

    // returns node's sum_and_count + number of subtrees's that are equal to their average
    var get_result func(node *TreeNode) (Sum_And_Count, int);
    get_result = func(node *TreeNode) (Sum_And_Count, int) {
        if node == nil { return Sum_And_Count{}, 0; }

        left_sum,  left_result  := get_result(node.Left);
        right_sum, right_result := get_result(node.Right);
        result_average := Sum_And_Count{
            sum  : left_sum.sum   + right_sum.sum   + node.Val,
            count: left_sum.count + right_sum.count + 1,
        };

        correct_count := left_result + right_result;
        if node.Val == result_average.sum / result_average.count {
            correct_count += 1;
        }

        return result_average, correct_count;
    }

    _, correct_count := get_result(root);
    return correct_count;
}
