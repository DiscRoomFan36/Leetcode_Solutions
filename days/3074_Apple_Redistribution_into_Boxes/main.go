package main

import "slices"

func minimumBoxes(apple []int, capacity []int) int {
    num_total_apples := 0
    for _, n := range apple { num_total_apples += n }

    slices.Sort(capacity)

    running_sum := 0
    // loop backwards
    for i := range capacity {
        // loop over backwards.
        cap := capacity[len(capacity) - i - 1]

        running_sum += cap
        if num_total_apples <= running_sum { return i+1 }
    }
    panic("not enough boxes to hold all apples")
}