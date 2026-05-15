package main

import "math"

func findMin(nums []int) int {
    n := len(nums);
    if n == 0 { panic("not possible"); }

    // this is both a "base case" and would make a faster check if n was small.
    if n <= 2 {
        smallest := math.MaxInt;
        for _, num := range nums {
            smallest = min(smallest, num)
        }
        return smallest;
    }

    // see if the array was rotated at all.
    if nums[0] < nums[n-1] {
        return nums[0]; // it wasnt rotated at all.
    }

    if nums[0] < nums[n/2] {
        // its in the upper half.
        return findMin(nums[n/2:]);
    } else {
        // its in the lower half.
        return findMin(nums[:n/2 + 1]);
    }
}
