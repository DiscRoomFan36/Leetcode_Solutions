package main

func scoreOfParentheses(s string) int {
    score := 0;
    stack := make([]int, 0);
    for _, c := range s {
        if c == '(' {
            Append(&stack, 0);
        } else {
            this_score := Pop(&stack);
            if this_score == 0 { this_score = 1; }

            if len(stack) > 0 {
                stack[len(stack)-1] += 2 * this_score;
            } else {
                score += this_score;
            }
        }
    }
    return score;
}

func Append[T any](slice *[]T, items ...T) *T {
    *slice = append(*slice, items...);
    return &(*slice)[len(*slice)-1];
}

func Pop[T any](slice *[]T) T {
    item  := (*slice)[ len(*slice)-1];
    *slice = (*slice)[:len(*slice)-1];
    return item;
}
