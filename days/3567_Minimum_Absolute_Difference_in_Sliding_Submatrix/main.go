package main

import (
	"math"
	"slices"
)

func minAbsDiff(grid [][]int, k int) [][]int {
    if len(grid) < k { return [][]int{}; }
    m, n := len(grid), len(grid[0]);

    result := make_grid[int](n-k+1, m-k+1);

    for y := range m-k+1 {
        for x := range n-k+1 {

            // find min distance
            numbers_in_submatrix := make([]int, 0, k*k);
            for j := range k {
                for i := range k {
                    cell := grid[y+j][x+i];
                    Append(&numbers_in_submatrix, cell);
                }
            }


            // Sort numbers for faster search.
            slices.Sort(numbers_in_submatrix);

            const INF = math.MaxInt;

            min_sum := INF;
            for i := range len(numbers_in_submatrix)-1 {
                a := numbers_in_submatrix[i  ];
                b := numbers_in_submatrix[i+1];
                if a == b { continue; }

                min_sum = min(min_sum, b - a);
            }

            if min_sum != INF {
                result[y][x] = min_sum;
            }
        }
    }

    return result;
}

func Append[T any](slice *[]T, items ...T) {
	*slice = append(*slice, items...);
}

// makes a grid, with default values. i hope go can turn this into something cool.
//
// might be better to make([]T, m*n) then grab some slices.
func make_grid[T any](n, m int) [][]T {
    grid := make([][]T, m);
    for j := range m { grid[j] = make([]T, n); }
    return grid;
}
