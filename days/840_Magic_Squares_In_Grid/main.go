package main

import "math/bits"

const MAGIC_SQUARE_SIZE = 3;

func numMagicSquaresInside(grid [][]int) int {
    if len(grid) == 0 { return 0; }
    m, n := len(grid), len(grid[0]);

    num_magic_squares := 0;
    for j := range m - MAGIC_SQUARE_SIZE+1 {
        for i := range n - MAGIC_SQUARE_SIZE+1 {
            // could do something faster like making a prefix sum,
            // but since this is a leetcode question, it expects
            // people in python to also compleate the question
            // in a reasonable time, so i dont have to worry about performace.
            if is_magic_square(grid, i, j) {
                num_magic_squares += 1;
            }
        }
    }
    return num_magic_squares;
}

func is_magic_square(grid [][]int, x, y int) bool {
    dr_diag := grid[y  ][x  ] + grid[y+1][x+1] + grid[y+2][x+2];
    ul_diag := grid[y+2][x  ] + grid[y+1][x+1] + grid[y  ][x+2];
    if dr_diag != ul_diag { return false; }

    correct_sum := ul_diag;

    // used to check for distinct digits.
    bit_field := uint(0);

    col_sums := [3]int{};
    for j := range MAGIC_SQUARE_SIZE {
        row_sum := 0;
        for i := range MAGIC_SQUARE_SIZE {
            cell := grid[y+j][x+i];
            // only numbers in range [1..9]
            if !(1 <= cell && cell <= 9) { return false; }

            bit_field |= 1 << cell;

            row_sum     += cell;
            col_sums[i] += cell;
        }
        if row_sum != correct_sum { return false; }
    }
    for _, sum := range col_sums {
        if sum != correct_sum { return false; }
    }

    // check that all cells were distinct.
    if bits.OnesCount(bit_field) != MAGIC_SQUARE_SIZE*MAGIC_SQUARE_SIZE {
        return false;
    }

    return true;
}
