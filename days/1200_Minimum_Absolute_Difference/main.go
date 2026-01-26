package main

import "slices"

func minimumAbsDifference(arr []int) [][]int {
    slices.Sort(arr)

    min_diff := 999999999999
    result := make([][]int, 0)

    for i := range len(arr)-1 {
        a, b := arr[i], arr[i+1]

        diff := b - a
        if diff < min_diff {
            min_diff = diff
            result = result[:0] // clear array
        }

        if diff == min_diff {
            result = append(result, []int{a, b})
        }
    }

    return result
}