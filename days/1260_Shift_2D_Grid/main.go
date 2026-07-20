package main

func shiftGrid(grid [][]int, k int) [][]int {
    if len(grid) == 0 { panic("what?"); }
    m, n := len(grid), len(grid[0]);

    if k % (n*m) == 0 { return grid; }

    result := make_grid[int](n, m);

    for i := range n*m {
        j := (i + k) % (n*m);
        result[j/n][j%n] = grid[i/n][i%n];
    }

    return result;
}

// makes a grid, with default values. i hope go can turn this into something cool.
//
// might be better to make([]T, m*n) then grab some slices.
func make_grid[T any](n, m int) [][]T {
    grid := make([][]T, m);
    for j := range m { grid[j] = make([]T, n); }
    return grid;
}
