package main

import (
	"math"
	"slices"
)

func search(nums []int, target int) int {
    min_index := findMinIndex(nums);
    // fmt.Println(min_index);

    left, right := nums[:min_index], nums[min_index:];

    // could do this smarter, i dont care.
    index_left,  found_left  := slices.BinarySearch(left,  target);
    if found_left  { return index_left; }

    index_right, found_right := slices.BinarySearch(right, target);
    if found_right { return index_right + len(left); } // remember to shift the index here

    return -1;
}

func findMinIndex(nums []int) int {

    var recur func(nums []int, low_index, high_index int) int;
    recur = func(nums []int, low_index, high_index int) int {
        n := len(nums);
        if n == 0 { panic("not possible"); }

        // this is both a "base case" and would make a faster check if n was small.
        if n <= 2 {
            smallest := math.MaxInt;
            smallest_index := -1;
            for i, num := range nums {
                if num < smallest {
                    smallest = num;
                    smallest_index = i;
                }
            }
            return smallest_index + low_index;
        }

        // see if the array was rotated at all.
        if nums[0] < nums[n-1] {
            return low_index;
        }

        if nums[0] < nums[n/2] {
            // its in the upper half.
            return recur(nums[n/2:], low_index+n/2, high_index);
        } else {
            // its in the lower half.
            return recur(nums[:n/2 + 1], low_index, low_index + n/2);
        }
    }

    return recur(nums, 0, len(nums)-1);
}