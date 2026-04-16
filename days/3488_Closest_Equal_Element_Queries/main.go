package main

import "slices"

func solveQueries(nums []int, queries []int) []int {
    number_positions := make(map[int][]int);
    for i, n := range nums {
        number_positions[n] = append(number_positions[n], i);
    }

    result := make([]int, len(queries));
    for i, q := range queries {
        num := nums[q];
        indexes := number_positions[num];
        if len(indexes) == 1 {
            result[i] = -1;
        } else {
            index_position, found := slices.BinarySearch(indexes, q);
            if found != true { panic("what"); }

            left__nums_index := indexes[Proper_Mod(index_position-1, len(indexes))];
            right_nums_index := indexes[Proper_Mod(index_position+1, len(indexes))];

            best_dist_left  := circle_index_distance(q, left__nums_index, len(nums));
            best_dist_right := circle_index_distance(q, right_nums_index, len(nums));

            result[i] = min(best_dist_left, best_dist_right);
        }
    }

    return result;
}

func Abs(x int) int {
    if x < 0 { x *= -1; }
    return x;
}

func Proper_Mod(x, y int) int {
    return ((x % y) + y) % y;
}

func circle_index_distance(a, b int, array_len int) int {
    if a <= b {
        return min(b - a, (a + array_len) - b);
    } else {
        return min(a - b, (b + array_len) - a);
    }
}
