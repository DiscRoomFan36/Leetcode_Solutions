package main

func minJumps(arr []int) int {
    // num to other num teleports.
    num_to_indexes := make(map[int][]int);
    for i, num := range arr {
        num_to_indexes[num] = append(num_to_indexes[num], i);
    }

    been_to := make([]bool, len(arr));
    // so we dont have to loop though num_to_indexes tones of times.
    number_checked := make(map[int]bool);

    curr_stack := make([]int, 0);
    next_stack := make([]int, 0);

    steps := 0;

    add_new_pos := func(new_pos int) {
        if new_pos < 0         { return; }
        if new_pos >= len(arr) { return; }
        if been_to[new_pos]    { return; }
        been_to[new_pos] = true;

        Append(&next_stack, new_pos);
    }

    add_new_pos(0);

    for {
        curr_stack, next_stack = next_stack, curr_stack;
        next_stack = next_stack[:0];

        if len(curr_stack) == 0 { panic("what"); }

        for _, pos := range curr_stack {
            if pos == len(arr)-1 { return steps; }

            add_new_pos(pos-1);
            add_new_pos(pos+1);
            if !number_checked[arr[pos]] {
                number_checked[arr[pos]] = true;
                for _, other_pos := range num_to_indexes[arr[pos]] {
                    add_new_pos(other_pos);
                }
            }
        }

        steps += 1;
    }
}


func Append[T any](slice *[]T, items ...T) *T {
    *slice = append(*slice, items...);
    return &(*slice)[len(*slice)-1];
}
