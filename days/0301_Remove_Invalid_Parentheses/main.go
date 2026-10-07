package main

import "math"

func removeInvalidParentheses(s string) []string {
    
    min_removals := math.MaxInt;
    result_map := make(map[string]struct{});

    var recur func(index int, stack_count int, removals int, current_result []byte);
    recur = func(index int, stack_count int, removals int, current_result []byte) {
        if removals > min_removals { return; }

        if index == len(s) {
            // check its valid parens
            if stack_count != 0 { return; }

            if removals < min_removals {
                min_removals = removals;
                clear(result_map);
            }

            result_map[string(current_result)] = struct{}{};
            return;
        }

        c := s[index];
        if c == '(' {
            recur(index + 1, stack_count+1, removals,     append(current_result, c));
            recur(index + 1, stack_count,   removals + 1, current_result);

        } else if c == ')' {
            if stack_count > 0 {
                recur(index + 1, stack_count - 1, removals, append(current_result, c));
            }
            recur(index + 1, stack_count, removals + 1, current_result);
            
        } else {
            recur(index + 1, stack_count, removals, append(current_result, c));
        }
    }

    recur(0, 0, 0, []byte{});

    result := make([]string, 0, len(result_map));
    for k, _ := range result_map { result = append(result, k); }
    return result;
}
