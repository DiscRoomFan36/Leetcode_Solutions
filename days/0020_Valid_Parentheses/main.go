package main

func isValid(s string) bool {
    stack := make([]rune, 0);

    OPENING  := [...]rune{'(', '[', '{'};
    CLOSEING := [...]rune{')', ']', '}'};

    for _, c := range s {
        for _, open := range OPENING {
            if c == open { stack = append(stack, open); }
        }

        for i, close := range CLOSEING {
            if c == close {
                if len(stack) == 0 { return false; }
                last := Pop(&stack);
                if OPENING[i] != last { return false; }
            }
        }
    }

    return len(stack) == 0;
}

func Pop[T any](slice *[]T) T {
    item  := (*slice)[ len(*slice)-1];
    *slice = (*slice)[:len(*slice)-1];
    return item;
}
