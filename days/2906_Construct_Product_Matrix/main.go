package main

func constructProductMatrix(grid [][]int) [][]int {
    if len(grid) == 0 { return [][]int{}; }
    m, n := len(grid), len(grid[0]);

    const MOD = 12345;

    p := make_grid[int](n, m);

    suffix := make([]int, n*m + 1);
    prefix := make([]int, n*m + 1);
    suffix[n*m] = 1;
    prefix[  0] = 1;

    index_to_matrix := func(grid [][]int, index int) *int {
        return &grid[index / n][index % n];
    }

    for i := range n*m {
        j := n*m - i - 1;
 
        suffix[j  ] = (suffix[j+1] * (*index_to_matrix(grid, j))) % MOD;
        prefix[i+1] = (prefix[i  ] * (*index_to_matrix(grid, i))) % MOD;
    }

    for i := range n*m {
        *index_to_matrix(p, i) = (prefix[i] * suffix[i+1]) % MOD;
    }

    return p;
}

// makes a grid, with default values. i hope go can turn this into something cool.
//
// might be better to make([]T, m*n) then grab some slices.
func make_grid[T any](n, m int) [][]T {
    grid := make([][]T, m);
    for j := range m { grid[j] = make([]T, n); }
    return grid;
}
