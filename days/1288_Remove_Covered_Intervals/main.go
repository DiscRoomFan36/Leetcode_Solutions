package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Interval struct {
    low, high int;
};

// if b is covering a, return true.
func is_covered(a, b Interval) bool {
    return b.low <= a.low && a.high <= b.high;
}

func removeCoveredIntervals(_intervals [][]int) int {

	intervals := make([]Interval, len(_intervals))
	for i, interval := range _intervals {
        intervals[i] = Interval{low: interval[0], high: interval[1]};
	}

    slices.SortFunc(intervals,
        func (a, b Interval) int {
            a_size := a.high - a.low;
            b_size := b.high - b.low;
            // largest intervals first
            return cmp.Compare(b_size, a_size);
        },
    );

    fmt.Println(intervals);
    safe_intervals := make([]Interval, 0);

    for _, new_interval := range intervals {
        is_safe := true;
        for _, safe_interval := range safe_intervals {
            if is_covered(new_interval, safe_interval) {
                is_safe = false;
            }
        }
        if is_safe {
            safe_intervals = append(safe_intervals, new_interval);
        }
    }

    return len(safe_intervals);
}
