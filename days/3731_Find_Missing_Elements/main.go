package main

import "math"

func findMissingElements(nums []int) []int {
    smallest := math.MaxInt;
    largest  := math.MinInt;
    // surly it would be faster just to sort...
    set := make(map[int]struct{});
    for _, n := range nums {
        smallest = min(smallest, n);
        largest  = max(largest,  n);
        set[n] = struct{}{};
    }

    result := []int{};
    for i := smallest+1; i < largest; i++ {
        _, in_set := set[i];
        if !in_set { result = append(result, i); }
    }
    return result;
}
