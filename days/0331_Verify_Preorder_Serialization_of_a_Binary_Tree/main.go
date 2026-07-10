package main

import (
	"strings"
)

func isValidSerialization(preorder string) bool {
    // handling the rare case when there is no tree.
    if preorder == "#" { return true; }

    bool_stack := make([]bool, 0);

    for i, c := range strings.Split(preorder, ",") {
        if c == "#" {
            if len(bool_stack) == 0 { return false; }

            for len(bool_stack) != 0 {
                is_right := Pop(&bool_stack);
                if !is_right {
                    Append(&bool_stack, true);
                    break;
                }
                // continue pop'ing,
            }

        } else {
            if len(bool_stack) == 0 {
                if i != 0 { return false; }
            }
            Append(&bool_stack, false);
        }
    }

    if len(bool_stack) == 0 {
        return true;
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
