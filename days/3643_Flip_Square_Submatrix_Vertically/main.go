package main

func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
    for i := range k/2 {
        j := k-i-1;

        for l := range k {
            grid[x+i][y+l], grid[x+j][y+l] = grid[x+j][y+l], grid[x+i][y+l];
        }
    }

    return grid;
}