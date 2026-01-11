package main

import (
	"cmp"
	"slices"
)

func earliestFullBloom(plantTime []int, growTime []int) int {
	n := len(plantTime)
	if n == 0 {
		return 0
	}

	// 1 <= plantTime[i], growTime[i] <= 104
	type plant_stats struct{ pt, gt int32 }

	plants := make([]plant_stats, n)
	for i := range n {
		plants[i] = plant_stats{
			pt: int32(plantTime[i]),
			gt: int32(growTime[i]),
		}
	}

	slices.SortFunc(plants,
		func(a, b plant_stats) int {
			return cmp.Compare(a.gt, b.gt)
		},
	)

	max_time := int32(0)
	running_sum := int32(0)
	// reverse iterate
	for i := n - 1; i >= 0; i-- {
		running_sum += plants[i].pt
		time_to_sprout := running_sum + plants[i].gt
		max_time = max(max_time, time_to_sprout)
	}

	return int(max_time)
}
