package main

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	if n != len(grid[0]) {
		panic("only square grids, although it would be easy to correct for this.")
	}

	// find the highest building in every row and collum.
	row_max := make([]int, n)
	col_max := make([]int, n)

	for j := range n {
		for i := range n {
			x := grid[j][i]

			row_max[j] = max(row_max[j], x)
			col_max[i] = max(col_max[i], x)
		}
	}

	// make a new grid with the max height that can go there.
	result := 0
	for j := range n {
		for i := range n {
			x := grid[j][i]
			max_height_this_can_be := min(row_max[j], col_max[i])
			result += max_height_this_can_be - x
		}
	}
	return result
}
