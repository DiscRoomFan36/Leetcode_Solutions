package main

func maxProductPath(grid [][]int) int {
    if len(grid) == 0 { return -1; }
    m, n := len(grid), len(grid[0]);

    next_max_sums := make([]int, n);
    next_min_sums := make([]int, n);

    prev_max_sums := make([]int, n);
    prev_min_sums := make([]int, n);

    next_max_sums[0] = grid[0][0];
    next_min_sums[0] = grid[0][0];
    for i := range n-1 {
        next_max_sums[i+1] = next_max_sums[i] * grid[0][i+1];
        next_min_sums[i+1] = next_min_sums[i] * grid[0][i+1];
    }

    for j := 1; j < m; j++ {
        // swap buffers
        next_max_sums, prev_max_sums = prev_max_sums, next_max_sums;
        next_min_sums, prev_min_sums = prev_min_sums, next_min_sums;

        next_max_sums[0] = prev_max_sums[0] * grid[j][0];
        next_min_sums[0] = prev_min_sums[0] * grid[j][0];

        for i := 1; i < n; i++ {
            cell := grid[j][i];

            // choose best from all options
            a, b := next_max_sums[i-1] * cell, prev_max_sums[i] * cell;
            c, d := next_min_sums[i-1] * cell, prev_min_sums[i] * cell;

            next_max_sums[i] = max(a, b, c, d);
            next_min_sums[i] = min(a, b, c, d);
        }
    }

    const MOD = 1000000007;

    result := next_max_sums[n-1];
    if result < 0 { return -1; }
    return result % MOD;
}
