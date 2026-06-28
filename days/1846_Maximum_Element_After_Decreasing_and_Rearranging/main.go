package main

import "slices"

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
    slices.Sort(arr);

    best := 0;
    for _, n := range arr {
        if n > best {
            best += 1;
        }
    }
    return best;
}
