package main

import "math"

func findPrefixScore(nums []int) []int64 {
    n := len(nums);
    score_array := make([]int64, n);

    max_so_far := int64(math.MinInt);
    conversion_running_sum := int64(0);
    for i := range nums {
        num := int64(nums[i]);

        max_so_far = max(max_so_far, num);
        conversion_running_sum += max_so_far + num;

        score_array[i] = conversion_running_sum;
    }

    return score_array;
}