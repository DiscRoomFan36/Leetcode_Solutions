package main

func countSubmatrices(grid [][]int, k int) int {
    if len(grid) == 0 { return 0; }
    m, n := len(grid), len(grid[0]);

    if grid[0][0] > k { return 0; }

    prefix_sum_grid := make_grid[int](n+1, m+1);

    for j := range m {
        for i := range n {
            up      := prefix_sum_grid[j  ][i+1];
            left    := prefix_sum_grid[j+1][i  ];
            up_left := prefix_sum_grid[j  ][i  ];
            prefix_sum_grid[j+1][i+1] = (up + left - up_left) + grid[j][i];
        }
    }

    get_area_sum := func(x, y, w, h int) int {
        return prefix_sum_grid[y+h][x+w] - prefix_sum_grid[y+h][x] - prefix_sum_grid[y][x+w] + prefix_sum_grid[y][x];
    }

    result := 0;
    for j := range m {
        for i := range n {
            if get_area_sum(0, 0, i+1, j+1) <= k {
                result += 1;
            } else {
                // all number in grid are positive, can only get bigger
                break;
            }
        }
    }
    return result;
}

// makes a grid, with default values. i hope go can turn this into something cool.
//
// might be better to make([]T, m*n) then grab some slices.
func make_grid[T any](n, m int) [][]T {
    grid := make([][]T, m)
    for j := range m { grid[j] = make([]T, n) }
    return grid
}
