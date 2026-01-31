package main

func maxSideLength(mat [][]int, threshold int) int {
    if len(mat) == 0 { return 0 }
    m, n := len(mat), len(mat[0])

    // should be init to zero.
    prefix_sum_grid := make_grid[int](n+1, m+1)

    for j := range m {
        for i := range n {
            up      := prefix_sum_grid[j  ][i+1]
            left    := prefix_sum_grid[j+1][i  ]
            up_left := prefix_sum_grid[j  ][i  ]
            prefix_sum_grid[j+1][i+1] = (up + left - up_left) + mat[j][i]
        }
    }

    get_area_sum := func(x, y, size int) int {
        return prefix_sum_grid[y+size][x+size] - prefix_sum_grid[y+size][x] - prefix_sum_grid[y][x+size] + prefix_sum_grid[y][x]
    }

    // forgot golang had goto's for loops, nice.
    max_size := 0
    y_loop: for y := range m {
        // dont go out of bounds
        if y + max_size+1 >= m+1 { break y_loop } 
        x_loop: for x := range n {
            // dont go out of bounds
            if x + max_size+1 >= n+1 { break x_loop }

            // grow the grid until it passes the threshold, or goes out of bounds
            inner_loop: for {
                area_sum := get_area_sum(x, y, max_size+1)
                if area_sum > threshold { break inner_loop }

                max_size += 1

                // gotta check these two again
                if x + max_size+1 >= n+1 { break x_loop }
                if y + max_size+1 >= m+1 { break y_loop }
            }
        }
    }
    return max_size
}

// makes a grid, with default values. i hope go can turn this into something cool.
//
// might be better to make([]T, m*n) then grab some slices.
func make_grid[T any](n, m int) [][]T {
    grid := make([][]T, m)
    for j := range m { grid[j] = make([]T, n) }
    return grid
}