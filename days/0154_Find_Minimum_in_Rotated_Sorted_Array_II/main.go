package main

import "math"

func findMin(nums []int) int {
    smallest := math.MaxInt;
    for _, num := range nums {
        smallest = min(smallest, num);
    } 
    return smallest;
}