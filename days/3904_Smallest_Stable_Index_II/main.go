package main

import "math"

func firstStableIndex(nums []int, k int) int {
    running_min := make([]int, len(nums));
    running_min[len(nums)-1] = nums[len(nums)-1];
    for i := len(nums)-2; i >= 0; i-- {
        running_min[i] = min(nums[i], running_min[i+1]);
    }

    running_max := math.MinInt;
    for i, num := range nums {
        running_max = max(running_max, num);

        score := running_max - running_min[i];
        // fmt.Println("i:", i, "score:", score);

        if score <= k { return i; }
    }

    return -1;
}
