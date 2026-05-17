package main

func canReach(arr []int, start int) bool {
    been_to := make([]bool, len(arr));

    stack := make([]int, 0);

    add_new_pos := func(new_pos int) {
        if new_pos < 0         { return; }
        if new_pos >= len(arr) { return; }
        if been_to[new_pos]    { return; }
        been_to[new_pos] = true;
        // maybe check for arr[new_pos] == 0 here?
        //
        // kinda want jai's #expand `return true;

        Append(&stack, new_pos);
    };

    add_new_pos(start);

    for len(stack) > 0 {
        pos := Pop(&stack);
        if arr[pos] == 0 { return true; }

        add_new_pos(pos - arr[pos]);
        add_new_pos(pos + arr[pos]);
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
