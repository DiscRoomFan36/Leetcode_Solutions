package main

import "slices"

func minOperations(grid [][]int, x int) int {
    if len(grid) == 0 { return 0; }
    m, n := len(grid), len(grid[0]);

    array := make([]int, m*n);
    for j := range m {
        for i := range n {
            array[j*n+i] = grid[j][i];
        }
    }

    slices.Sort(array);

    middle := array[m*n / 2];

    result := 0;
    for _, num := range array {
        if num % x != middle % x { return -1; }
        result += Abs(middle - num) / x;
    }
    return result;
}

func Abs(x int) int {
    if x < 0 { x *= -1; }
    return x;
}