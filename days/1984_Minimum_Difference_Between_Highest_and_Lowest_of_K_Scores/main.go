package main

import "slices"

func minimumDifference(nums []int, k int) int {
    if k <= 1         { return 0 }
    if len(nums) == 0 { return 0 }

    slices.Sort(nums)

    best := 999999999999
    for i := range len(nums) - (k-1) {
        low := nums[i]
        // if k == 2, we need i+1
        high := nums[i+k-1]

        diff := high - low
        best = min(best, diff)
    }

    return best
}
