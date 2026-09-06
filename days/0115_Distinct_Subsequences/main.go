package main

func numDistinct(s string, t string) int {
    curr_dp := make([]int, len(t)+1);
    next_dp := make([]int, len(t)+1);

    curr_dp[0] = 1;

    for i := len(s)-1; i >= 0; i -= 1 {
        char := s[i];

        next_dp[0] = curr_dp[0];
        for j := range len(t) {
            // correct characters bubble along.
            next_dp[j+1] = curr_dp[j+1];
            if t[len(t)-j-1] == char {
                next_dp[j+1] += curr_dp[j];
            }
        }

        curr_dp, next_dp = next_dp, curr_dp;
    }

    return curr_dp[len(t)];
}
