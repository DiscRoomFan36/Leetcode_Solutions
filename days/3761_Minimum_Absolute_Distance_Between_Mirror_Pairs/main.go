package main

import "math"

func minMirrorPairDistance(nums []int) int {
    const INF = math.MaxInt;
    best_distance := INF;

    num_to_index := make(map[int]int);
    for i := len(nums)-1; i >= 0; i-- {
        num := nums[i];
        if index, found := num_to_index[reverse(num)]; found {
            best_distance = min(best_distance, index - i);
        }

        num_to_index[num] = i;
    }

    if best_distance == INF { return -1; }
    return best_distance;
}

func reverse(x int) int {
    result := 0;
    for x > 0 {
        result = result * 10 + x % 10;
        x /= 10;
    }
    return result;
}
