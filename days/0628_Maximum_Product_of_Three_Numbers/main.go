package main

import "slices"

func maximumProduct(nums []int) int {
    slices.Sort(nums);
    n := len(nums);
    return max(
        nums[  0] * nums[  1] * nums[n-1],
        nums[n-3] * nums[n-2] * nums[n-1],
    );
}
