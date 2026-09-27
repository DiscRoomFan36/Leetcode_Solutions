package main

func reverseParentheses(s string) string {
    result := make([]byte, 0);

    start_stack := make([]int, 0);
    for i := range len(s) {
        if s[i] == '(' {
            Append(&start_stack, len(result));

        } else if s[i] == ')' {
            if len(start_stack) == 0 { panic("unbalenced parentheses"); }
            start := Pop(&start_stack);

            // reverse previous section
            reverse_slice(result[start:]);
        } else {
            // just a normal character
            Append(&result, s[i]);
        }
    }

    return string(result);
}

func reverse_slice[T any](slice []T) {
    for i := 0; i < len(slice)/2; i += 1 {
        j := len(slice)-1-i;
        slice[i], slice[j] = slice[j], slice[i];
    }
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
