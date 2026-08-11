package main

func missingInteger(nums []int) int {
    prefix_sum := nums[0];
    for j := 1; j < len(nums); j++ {
        if nums[j] == nums[j - 1] + 1 {
            prefix_sum += nums[j];
        } else {
            break;
        }
    }

    set := make(map[int]struct{});
    for _, n := range nums {
        if n < prefix_sum { continue; }
        set[n] = struct{}{};
    }

    for result := prefix_sum; ; result += 1 {
        _, in_set := set[result];
        if !in_set { return result; }
    }
}
