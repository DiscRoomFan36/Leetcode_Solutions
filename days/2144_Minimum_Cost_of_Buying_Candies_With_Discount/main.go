package main

import "slices"

func minimumCost(cost []int) int {
    slices.Sort(cost);

    total_cost := 0;

    count := 0;
    for i := len(cost)-1; i >= 0; i-- {
        if count == 2 {
            count = 0;
            continue;
        }

        count += 1;
        total_cost += cost[i];
    }

    return total_cost;
}
