package main

import (
	"cmp"
	"math"
	"slices"
)

func shortestBeautifulSubstring(s string, k int) string {
    ones_positions := make([]int, 0);
    for i, c := range s {
        if c == '1' {
            ones_positions = append(ones_positions, i);
        }
    }
    // fmt.Println(ones_positions);

    // not enough to make a single string
    if len(ones_positions) < k { return ""; }

    // degenerate case
    if k == 1 { return "1"; }

    best_index  := 0;
    best_length := math.MaxInt;
    for i := range len(ones_positions) - k + 1 {
        start_index := ones_positions[i];
        end_index   := ones_positions[i + k - 1];

        length := end_index - start_index;
        if length < best_length {
            best_length = length
            best_index  = i;
        } else if length == best_length {
            // get lexicographically smallest
            compare := slices.CompareFunc(ones_positions[i:i+k], ones_positions[best_index:best_index+k],
                func(a, b int) int {
                    mapped_a := a - ones_positions[i];
                    mapped_b := b - ones_positions[best_index];
                    return cmp.Compare(mapped_a, mapped_b);
                },
            );
            // this may be backwards
            if compare > 0 { best_index = i; }
        }
    }

    start_index := ones_positions[best_index];
    end_index   := ones_positions[best_index + k - 1];

    return s[start_index:end_index+1];
}
