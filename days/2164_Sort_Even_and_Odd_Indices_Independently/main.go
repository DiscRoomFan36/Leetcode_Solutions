package main

import (
	"cmp"
	"slices"
)

func sortEvenOdd(nums []int) []int {
	// we *COULD* use the original array to prevent any allocations,
	// but this is an easy question, so i dont actually care.
	//
	// also if your allocator is fast enough, might be faster
	// than doing something more complicated,
	//
	// also this situation would never happen in real life.
	evens := make([]int, (len(nums)+1)/2)
	odds := make([]int, (len(nums))/2)

	for i := range nums {
		if i%2 == 0 {
			evens[i/2] = nums[i]
		} else {
			odds[i/2] = nums[i]
		}
	}

	slices.Sort(evens)
	// sort in reverse
	slices.SortFunc(odds, func(a, b int) int { return cmp.Compare(b, a) })

	for i := range nums {
		if i%2 == 0 {
			nums[i] = evens[i/2]
		} else {
			nums[i] = odds[i/2]
		}
	}

	return nums
}
