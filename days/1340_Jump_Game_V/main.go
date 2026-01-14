package main

import (
	"cmp"
	"slices"
)

func maxJumps(arr []int, d int) int {
	n := len(arr)

	// indexes into the max jumps array, sorted by the height at arr[i]
	sorted_indexes := make([]int, n)
	for i := range n {
		sorted_indexes[i] = i
	}

	slices.SortFunc(sorted_indexes,
		func(a, b int) int {
			return cmp.Compare(arr[a], arr[b])
		},
	)

	// how far this index can jump
	how_far := make([]int, n)

	max_distance := 0

	for i := range n {
		index := sorted_indexes[i]
		this_height := arr[index]

		max_number_of_jumps_can_do := 0
		for k := 1; k <= d; k++ {
			checking_index := index + k
			// dont go out of bounds
			if checking_index >= n {
				break
			}

			// we cant jump over this
			if arr[checking_index] >= this_height {
				break
			}

			// we can jump here
			max_number_of_jumps_can_do = max(max_number_of_jumps_can_do, how_far[checking_index])
		}

		for k := 1; k <= d; k++ {
			checking_index := index - k
			// dont go out of bounds
			if checking_index < 0 {
				break
			}

			// we cant jump over this
			if arr[checking_index] >= this_height {
				break
			}

			// we can jump here
			max_number_of_jumps_can_do = max(max_number_of_jumps_can_do, how_far[checking_index])
		}

		new_how_far := max_number_of_jumps_can_do + 1
		how_far[index] = new_how_far

		max_distance = max(max_distance, new_how_far)
	}

	return max_distance
}
