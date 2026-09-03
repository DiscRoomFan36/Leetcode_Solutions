package main

import "math"

func uniformArray(nums1 []int) bool {
    // either their all even,
    all_even := true;
    smallest_number := math.MaxInt;
    for _, n := range nums1 {
        smallest_number = min(smallest_number, n);

        if n % 2 == 1 { all_even = false; }
    }
    if all_even { return true; }

    // or we have to make them all odd.
    //
    // can only make an even number odd, when there is a smaller odd number,
    // so we check the smallest number's parity 
    return (smallest_number % 2) == 1;
}
