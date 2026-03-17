package main

import (
	"cmp"
	"slices"
)

func largestSubmatrix(matrix [][]int) int {
    if len(matrix) == 0 { return 0; }
    m, n := len(matrix), len(matrix[0]);

    // modify the matrix, give it col prefix sums.
    for j := 1; j < m; j++ {
        for i := range n {
            if matrix[j][i] != 0 {
                matrix[j][i] += matrix[j-1][i];
            }
        }
    }

    best_score := 0;
    for j := range m {
        // sort the rows, we dont care about orginal order anymore.
        //
        // biggest first.
        slices.SortFunc(matrix[j], func(a, b int) int { return cmp.Compare(b, a); });

        // the highest the rectangle can be.
        max_height := matrix[j][0];
        // the start value is a 1 wide stright up down.
        best_score = max(best_score, matrix[j][0])
        for i := range n {
            num := matrix[j][i];
            if num == 0 { break; }

            max_height = min(max_height, num);
            score := max_height * (i+1);
            best_score = max(best_score, score);
        }
    }
    return best_score;
}
