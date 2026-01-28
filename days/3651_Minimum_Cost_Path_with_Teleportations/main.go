package main

import (
	"cmp"
	"math"
	"slices"
)

func minCost(grid [][]int, _k int) int {
    if len(grid) == 0 { return 0 }

    m, n := len(grid), len(grid[0])

    const INF = math.MaxInt

    // dynamic programming,
    cost_grid := make([]int, m*n)
    for i := range cost_grid { cost_grid[i] = INF }

    type Point struct { x, y int}
    points := make([]Point, m*n)
    for j := range m {
        for i := range n {
            points[j*n + i] = Point{i, j}
        }
    }

    slices.SortFunc(points, 
        func(a, b Point) int {
            // sort in decending order.
            return -cmp.Compare(grid[a.y][a.x], grid[b.y][b.x])
        },
    )

    for range _k+1 {

        { // handle teleports.
            min_cost := INF
            i := 0
            j := 0

            for i < len(points) {
                // find the first difference thing.
                for {
                    min_cost = min(min_cost, cost_grid[points[i].y*n + points[i].x])
                    // so we dont run of the edge
                    if i + 1 >= len(points) { break }
                    // get all similar numbers
                    if grid[points[i].y][points[i].x] != grid[points[i + 1].y][points[i + 1].x] { break }
                    // else continue.
                    i += 1
                }

                for l := j; l <= i; l++ {
                    cost_grid[points[l].y*n + points[l].x] = min_cost
                }

                // goto next point.
                j = i + 1
                i = i + 1
            }
        }

        for y := range m {
            for x := range n {
                // dont do anything to the start cell, its always zero
                if y == 0 && x == 0 {
                    cost_grid[y*n + x] = 0
                    continue
                }

                this_grid := grid[y][x]

                cell_result := cost_grid[y*n + x]

                if x > 0 { cell_result = min(cell_result, cost_grid[(y  )*n + (x-1)] + this_grid) }
                if y > 0 { cell_result = min(cell_result, cost_grid[(y-1)*n + (x  )] + this_grid) }

                cost_grid[y*n + x] = cell_result
            }
        }
    }

    return cost_grid[(m-1)*n + (n-1)]
}