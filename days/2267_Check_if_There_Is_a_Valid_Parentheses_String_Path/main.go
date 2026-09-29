package main

func hasValidPath(grid [][]byte) bool {
    if len(grid) == 0 { panic("invalid leetcode input"); }
    m, n := len(grid), len(grid[0]);

    // the output string length is odd, cannot be valid.
    if (m + n - 1) % 2 != 0 { return false; }
    // wrong starting char.
    if grid[  0][  0] == ')' { return false; }
    // wrong ending char
    if grid[m-1][n-1] == '(' { return false; }

    type State struct { x, y, depth int; }

    state_set := make(map[State]struct{});

    stack := make([]State, 0);
    Append(&stack, State{0, 0, 1});

    for len(stack) > 0 {
        state := Pop(&stack);
        // if this the end tile
        if state.x == n-1 && state.y == m-1 {
            // check if we won
            if state.depth == 0 { return true; }
            // else just continue.
            continue;
        }

        add_new_state := func(delta_x, delta_y int) {
            new_state := State{
                state.x + delta_x,
                state.y + delta_y,
                state.depth,
            };
            if new_state.x >= n { return; }
            if new_state.y >= m { return; }

            if grid[new_state.y][new_state.x] == '(' {
                new_state.depth += 1;
            } else {
                if new_state.depth == 0 { return; }
                new_state.depth -= 1;
            }

            // check if the paren depth is too big to close before we reach the end
            left_to_go := (n - 1 - new_state.x) + (m - 1 - new_state.y);
            if new_state.depth > left_to_go { return; }

            _, found := state_set[new_state];
            if found { return; }
            state_set[new_state] = struct{}{};

            Append(&stack, new_state);
        }

        add_new_state(1, 0);
        add_new_state(0, 1);
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
