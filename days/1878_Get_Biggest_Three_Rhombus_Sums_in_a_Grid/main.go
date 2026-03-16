package main

import (
	"cmp"
	"slices"
)

func getBiggestThree(grid [][]int) []int {
    if len(grid) == 0 { return []int{}; }
    m, n := len(grid), len(grid[0]);

    all_sums := make([]int, 0);

    for j := range m {
        for i := range n {
            // [j][i] is the top of the rombas

            // size of the rect we making.
            for size := range min(n, m)/2 + 1 {

                // bounds checking
                if i + size >= n { break; }
                if i - size <  0 { break; }
                // bottom tip.
                if j + size*2 >= m { break; }

                sum := grid[j][i];
                // top half.
                for k := 1; k <= size; k++ {
                    sum += grid[j+k][i-k];
                    sum += grid[j+k][i+k];
                }
                // bottom half
                for k := 1; k < size; k++ {
                    sum += grid[j+size+k][i+size-k];
                    sum += grid[j+size+k][i-size+k];
                }
                if size != 0 {
                    sum += grid[j+size*2][i];
                }

                Append(&all_sums, sum);
            }
        }
    }

    // would be faster to just keep track of the top 3, but i dont like golangs heap library.
    slices.SortFunc(all_sums, func(a, b int) int { return cmp.Compare(b, a); });

    const NUM_MAX_SUMS = 3;
    best := make([]int, 0);
    // dedup
    for i := range all_sums {
        if i == 0 || all_sums[i] != all_sums[i-1] {
            Append(&best, all_sums[i]);
        }
        if len(best) >= NUM_MAX_SUMS { break; }
    }

    return best;
}

func Append[T any](slice *[]T, items ...T) {
	*slice = append(*slice, items...);
}
