package main

func numberOfSets(n int, k int) int {
    const MOD = 1000000007;
    dp := make([]int, n);
    prefix_sum := make([]int, n+1);

    for j := range n {
        dp[j] = 1;
        prefix_sum[j+1] = (prefix_sum[j] + dp[j]) % MOD;
    }

    for i := 1; i <= k; i += 1 {
        dp[0] = 0;
        for j := 1; j < n; j += 1 {
            dp[j] = (dp[j-1] + prefix_sum[j]) % MOD;
        }
        for j := 0; j < n; j += 1 {
            prefix_sum[j+1] = (prefix_sum[j] + dp[j]) % MOD;
        }
    }

    return dp[n-1];
}
