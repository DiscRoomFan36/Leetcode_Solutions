package main

func generateParenthesis(n int) []string {
    current := make(map[string]struct{});
    current["()"] = struct{}{};
    next    := make(map[string]struct{});
    for range n-1 {
        clear(next);
        for s, _ := range current {
            next["()" + s       ] = struct{}{};
            next["("  + s +  ")"] = struct{}{};
            next[       s + "()"] = struct{}{};
        }

        current, next = next, current;
    }

    result := make([]string, 0);
    for k, _ := range current {
        result = append(result, k);
    }
    return result;
}
