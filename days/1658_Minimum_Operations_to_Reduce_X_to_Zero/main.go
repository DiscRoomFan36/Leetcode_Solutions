package main

import "math"

func minOperations(nums []int, x int) int {
    const INF = math.MaxInt;

    left := 0;
    running := 0;
    current_length := 0;
    for ; left < len(nums); left += 1 {
        current_length += 1;
        running += nums[left];
        if running >= x { break; }
    }

    // couldnt sum to it anyway;
    if running < x { return -1; }

    best_length := INF;
    if running == x { best_length = current_length; }

    right := len(nums)-1;
    for left >= 0 {
        running -= nums[left];
        left -= 1;
        current_length -= 1;
        for running < x {
            running += nums[right];
            right -= 1;
            current_length += 1;
        }

        if running == x {
            best_length = min(best_length, current_length)
        }
    } 

    if best_length == INF { return -1; }
    return best_length;
}
