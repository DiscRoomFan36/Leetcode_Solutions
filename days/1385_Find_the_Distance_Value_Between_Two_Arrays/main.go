package main

import "slices"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
    slices.Sort(arr2);

    dist := 0;
    for _, a := range arr1 {
        index, _ := slices.BinarySearch(arr2, a);
        if index > 0         && Abs(a - arr2[index-1]) <= d { continue; }
        if index < len(arr2) && Abs(a - arr2[index  ]) <= d { continue; }
        dist += 1;
    }
    return dist;
}

func Abs(x int) int {
    if x < 0 { x = -x; }
    return x;
}
