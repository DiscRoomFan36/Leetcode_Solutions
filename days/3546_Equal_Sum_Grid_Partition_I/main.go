package main

func canPartitionGrid(grid [][]int) bool {
    if len(grid) == 0 { return false; }
    m, n := len(grid), len(grid[0]);

    prefix_sum_grid := make_grid[int](n+1, m+1);

    for j := range m {
        for i := range n {
            up      := prefix_sum_grid[j  ][i+1];
            left    := prefix_sum_grid[j+1][i  ];
            up_left := prefix_sum_grid[j  ][i  ];
            prefix_sum_grid[j+1][i+1] = (up + left - up_left) + grid[j][i];
        }
    }

    // includes the two end points.
    get_area_sum := func(x1, y1, x2, y2 int) int {
        w := x2 - x1 + 1;
        h := y2 - y1 + 1;
        return prefix_sum_grid[y1+h][x1+w] - prefix_sum_grid[y1+h][x1] - prefix_sum_grid[y1][x1+w] + prefix_sum_grid[y1][x1];
    }

    for i := range n {
        sum_1 := get_area_sum(0, 0,     i, m-1);
        sum_2 := get_area_sum(i+1, 0, n-1, m-1);
        if sum_1 == sum_2 { return true; }
    }
    for j := range m {
        sum_1 := get_area_sum(0, 0,   n-1,   j);
        sum_2 := get_area_sum(0, j+1, n-1, m-1);
        if sum_1 == sum_2 { return true; }
    }

    return false;
}

// makes a grid, with default values. i hope go can turn this into something cool.
//
// might be better to make([]T, m*n) then grab some slices.
func make_grid[T any](n, m int) [][]T {
    grid := make([][]T, m)
    for j := range m { grid[j] = make([]T, n) }
    return grid
}
