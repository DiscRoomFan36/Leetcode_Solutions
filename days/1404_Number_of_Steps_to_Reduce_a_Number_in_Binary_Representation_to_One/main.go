package main

import "strings"

func numSteps(s string) int {

    last_one_in_string := strings.LastIndexByte(s, '1');
    if last_one_in_string == 0 {
        // ony one set bit, just shift until we hit the end.
        return len(s)-1;
    }

    num_zeros_in_tail := len(s) - last_one_in_string - 1;
    without_tail := s[:last_one_in_string+1];

    num_zeros := strings.Count(without_tail, "0");
    num_ones  := len(without_tail) - num_zeros;

    return num_zeros_in_tail + num_zeros*2 + num_ones + 1;
}