package main

func maxDepth(s string) int {
    max_depth := 0;
    stack := 0;
    for _, c := range s {
        if c == '(' { stack += 1; }
        if c == ')' { stack -= 1; }
        max_depth = max(max_depth, stack);
    }
    return max_depth;
}
