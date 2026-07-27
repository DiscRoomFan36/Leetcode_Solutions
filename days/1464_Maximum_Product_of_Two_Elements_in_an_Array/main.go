package main

import "slices"

func maxProduct(nums []int) int {
    n := len(nums);
    slices.Sort(nums);
    multiply := func(a, b int) int {
        return (a-1) * (b-1);
    }
    return max(
        multiply(nums[  0], nums[  1]),
        multiply(nums[n-2], nums[n-1]),
    );
}
