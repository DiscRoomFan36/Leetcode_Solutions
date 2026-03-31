package main

import "math"

func maxMatrixSum(matrix [][]int) int64 {
    if len(matrix) == 0 { return 0; }
    m, n := len(matrix), len(matrix[0]);

    total_if_all_were_positive := 0;
    num_negative_numbers := 0;
    smallest_absolute_number := math.MaxInt;

    was_a_zero := false;

    for j := range m {
        for i := range n {
            cell := matrix[j][i];
            total_if_all_were_positive += abs(cell);
            smallest_absolute_number = min(smallest_absolute_number, abs(cell));

            if cell == 0 { was_a_zero = true; }
            if cell <  0 { num_negative_numbers += 1; }
        }
    }

    if was_a_zero || (num_negative_numbers % 2 == 0) {
        return int64(total_if_all_were_positive);
    }

    // just remove the smallest absolute number twice, (because we allready counted it.)
    return int64(total_if_all_were_positive + (-smallest_absolute_number)*2);
}

func abs(x int) int {
    if x < 0 { x *= -1; }
    return x;
}
