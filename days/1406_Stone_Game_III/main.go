package main

import "math"

func stoneGameIII(stoneValue []int) string {
    const INF = math.MaxInt;

    n := len(stoneValue);

    dp := make([]int, n+1);
    for i := range dp { dp[i] = -INF; }

    var recur func(i int) int;
    recur = func(i int) int {
        // base case, the game is over.
        if i >= n { return 0; }

        // dynamic programming,
        if dp[i] == -INF {
            // firgure out the best move

            // take 1 rock
            best_score := stoneValue[i] - recur(i + 1);
            // take 2 rocks
            if i + 1 < n {
                best_score = max(best_score, stoneValue[i] + stoneValue[i+1] - recur(i + 2));
            }
            // take 3 rocks
            if i + 2 < n {
                best_score = max(best_score, stoneValue[i] + stoneValue[i+1] + stoneValue[i+2] - recur(i + 3));
            }
            
            dp[i] = best_score;
        }


        return dp[i];
    };

    best_alice_score := recur(0);
    if best_alice_score > 0 {
        return "Alice";
    } else if best_alice_score == 0 {
        return "Tie";
    } else {
        return "Bob";
    }
}
