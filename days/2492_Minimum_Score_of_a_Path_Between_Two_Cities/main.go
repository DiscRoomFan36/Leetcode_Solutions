package main

import "math"

func minScore(n int, roads [][]int) int {
    // flood find.
    type Path_And_Dist struct {
        path, dist int;
    };
    num_to_roads := make(map[int][]Path_And_Dist);
    for _, road := range roads {
        a, b, distance := road[0], road[1], road[2];
        num_to_roads[a] = append(num_to_roads[a], Path_And_Dist{b, distance});
        num_to_roads[b] = append(num_to_roads[b], Path_And_Dist{a, distance});
    }

    been_to := make(map[int]bool);
    stack := []int{1}; // 1 is starting point.
    min_dist := math.MaxInt;
    for len(stack) > 0 {
        node := Pop(&stack);
        if been_to[node] { continue; }
        been_to[node] = true;

        for _, path_and_dist := range num_to_roads[node] {
            min_dist = min(min_dist, path_and_dist.dist);
            Append(&stack, path_and_dist.path);
        }
    }

    return min_dist;
}

func Append[T any](slice *[]T, items ...T) *T {
    *slice = append(*slice, items...);
    return &(*slice)[len(*slice)-1];
}

func Pop[T any](slice *[]T) T {
    item  := (*slice)[ len(*slice)-1];
    *slice = (*slice)[:len(*slice)-1];
    return item;
}
