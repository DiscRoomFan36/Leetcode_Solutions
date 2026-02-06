package main

import "slices"

func minRemoval(nums []int, k int) int {
    slices.Sort(nums)

    length := 1
    // move forwards along the array
    for start := 0; start < len(nums); start++ {
        // extend if possible
        for {
            if start+length >= len(nums) { break }
            if nums[start]*k < nums[start + length] { break }
            length += 1
        }
    }

    return len(nums) - length
}