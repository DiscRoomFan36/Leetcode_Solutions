package main

func longestSubsequence(nums []int) int {
    seen_non_zero := false;
    xor := 0;
    for _, n := range nums {
        if n != 0 {
            seen_non_zero = true;
        }
        xor ^= n;
    }

    if xor != 0      { return len(nums); }
    if seen_non_zero { return len(nums)-1; }
    return 0;
}
