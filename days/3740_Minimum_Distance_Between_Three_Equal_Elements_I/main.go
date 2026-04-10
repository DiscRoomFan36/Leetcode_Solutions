package main

import "math"

func minimumDistance(nums []int) int {
    const INF = math.MaxInt;

    type Index_Tracking struct {
        last_index int;
        last_last_index int;
    }

    prev_map := make(map[int]Index_Tracking);

    min_distance := INF;
    for this_index, num := range nums {
        // look at prev instances of num,

        index_tracking, found := prev_map[num];

        if !found {
            // initialize the struct.
            index_tracking.last_index      = -1;
            index_tracking.last_last_index = -1;
        }

        if index_tracking.last_last_index != -1 {
            // seen enough index's
            i, j, k := this_index, index_tracking.last_index, index_tracking.last_last_index;
            distance := Abs(i - j) + Abs(j - k) + Abs(k - i);
            min_distance = min(min_distance, distance);
        }

        // rotate and return to map
        index_tracking.last_last_index = index_tracking.last_index;
        index_tracking.last_index      = this_index;
        prev_map[num] = index_tracking;
    }

    if min_distance == INF { return -1; }
    return min_distance;
}

func Abs(x int) int {
    if x < 0 { x *= -1; }
    return x;
}
