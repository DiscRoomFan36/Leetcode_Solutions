
package main

func numberOfSubmatrices(grid [][]byte) int {
    if len(grid) == 0 { return 0 }

    total := 0

    num_x_in_column := make([]int, len(grid[0]))
    num_y_in_column := make([]int, len(grid[0]))

    for j := range grid {
        num_x_this_row := 0
        num_y_this_row := 0

        for i := range grid[j] {
            num_x_in_column[i] += num_x_this_row
            num_y_in_column[i] += num_y_this_row

            if grid[j][i] == 'X' {
                num_x_in_column[i] += 1
                num_x_this_row     += 1
            }
            if grid[j][i] == 'Y' {
                num_y_in_column[i] += 1
                num_y_this_row     += 1
            }

            num_xs := num_x_in_column[i]
            num_ys := num_y_in_column[i]

            if num_xs > 0 && num_xs == num_ys { total += 1 }
        }
    }

    return total
}