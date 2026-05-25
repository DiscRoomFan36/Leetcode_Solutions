package main

func canReach(s string, minJump int, maxJump int) bool {
    if s[0] != '0' { panic("what"); }
    // would be kinda weird.
    if s[len(s)-1] != '0' { return false; }

    jumpable_positions := make([]int, 0);
    for i, c := range s {
        if c == '0' {
            Append(&jumpable_positions, i);
        }
    }

    been_to := make([]bool, len(jumpable_positions));

    // contains a jumpable_index
    jumpable_index_stack := make([]int, 0);
    Append(&jumpable_index_stack, 0);
    been_to[0] = true;

    for len(jumpable_index_stack) > 0 {
        jumpable_index := Pop(&jumpable_index_stack);
        if jumpable_index == len(jumpable_positions)-1 { return true; }

        pos := jumpable_positions[jumpable_index];
        for i := jumpable_index+1; i < len(jumpable_positions); i++ {
            next_pos := jumpable_positions[i];
            if next_pos - pos < minJump { continue; }
            if next_pos - pos > maxJump { break; }

            if been_to[i] { continue; }
            been_to[i] = true;

            Append(&jumpable_index_stack, i);
        }
    }

    return false;
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
