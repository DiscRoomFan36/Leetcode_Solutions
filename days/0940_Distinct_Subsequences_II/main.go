package main

func distinctSubseqII(s string) int {
    dp := make([]int, len(s)+1);
    dp[0] = 1;
    last := [26]int{};
    for i := range last { last[i] = -1; }

    const MOD = 1000000007;

    for i, c := range s {
        dp[i+1] = (dp[i] * 2) % MOD;
        if last[c - 'a'] >= 0 {
            dp[i+1] -= dp[last[c - 'a']];
        }
        dp[i+1] %= MOD;
        last[c - 'a'] = i;
    }
    return ((dp[len(s)] - 1) % MOD + MOD) % MOD;
}
