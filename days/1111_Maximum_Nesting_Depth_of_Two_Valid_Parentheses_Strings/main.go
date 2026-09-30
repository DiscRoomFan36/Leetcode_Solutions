package main

func maxDepthAfterSplit(seq string) []int {
    result := make([]int, len(seq));
    depth := 0;
    for i, c := range seq {
        if c == '(' { depth += 1; }
        result[i] = depth % 2;
        if c == ')' { depth -= 1; }
    }
    return result;
}
