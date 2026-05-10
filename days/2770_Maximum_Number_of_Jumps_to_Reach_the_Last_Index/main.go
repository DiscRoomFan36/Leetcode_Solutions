package main

func maximumJumps(nums []int, target int) int {
    n := len(nums);

    dp := make([]int, n);

    for i := n-2; i >= 0; i-- {
        maximum_jumps := -1;
        for j := i+1; j < n; j++ {
            if dp[j] == -1 { continue; }
            if !Is_Between(nums[j] - nums[i], -target, target) { continue; }

            maximum_jumps = max(maximum_jumps, dp[j]+1);
        }

        dp[i] = maximum_jumps;
    }

    result := dp[0];
    return result;
}

func Is_Between(x int, low, high int) bool {
    return low <= x && x <= high;
}
