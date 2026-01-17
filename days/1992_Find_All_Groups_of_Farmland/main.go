package main

func findFarmland(land [][]int) [][]int {
	m, n := len(land), len(land[0])

	result := make([][]int, 0)

	for j := range m {
		for i := range n {
			// check if we need to do this one.
			if land[j][i] == 0 {
				continue
			}

			// if the previous one is on, it was handled previously
			if i > 0 && land[j][i-1] != 0 {
				continue
			}
			if j > 0 && land[j-1][i] != 0 {
				continue
			}

			width := 0
			for k := i + 1; k < n; k++ {
				if land[j][k] == 0 {
					break
				}
				width += 1
			}

			height := 0
			for k := j + 1; k < m; k++ {
				if land[k][i] == 0 {
					break
				}
				height += 1
			}

			result = append(result, []int{j, i, j + height, i + width})

			// skip the width that we know is done.
			// someone who likes go might hate me for this
			i += width + 1
		}
	}

	return result
}
