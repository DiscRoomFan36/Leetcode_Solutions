package main

func stoneGameVIII(stones []int) int {
    n := len(stones);

    running_sum := make([]int, n);
    running_sum[0] = stones[0];
    for i := 1; i < n; i += 1 {
        running_sum[i] = running_sum[i-1] + stones[i];
    }

    dp := make([]int, n);
    dp[n-1] = running_sum[n-1];
    for i := len(stones) - 2; i >= 1; i -= 1 {
        dp[i] = max(dp[i+1], running_sum[i]-dp[i+1]);
    }
    return dp[1];
}
