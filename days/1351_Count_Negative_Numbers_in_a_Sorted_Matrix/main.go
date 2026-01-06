
package main

func countNegatives(grid [][]int) int {
    if len(grid) == 0 { return 0 }
    m, n := len(grid), len(grid[0])

    total := 0

    j := 0
    for i := n-1; i >= 0; i -= 1 {
        for j < m && grid[j][i] >= 0 { j += 1 }
        if j >= m { break }

        total += m - j
    }

    return total
}
