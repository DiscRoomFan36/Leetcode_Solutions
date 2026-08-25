package main

func missingMultiple(nums []int, k int) int {
    nums_as_map := make(map[int]struct{});
    for _, n := range nums {
        nums_as_map[n] = struct{}{};
    }

    for result := k; ; result += k {
        _, ok := nums_as_map[result];
        if !ok { return result; }
    }
}
