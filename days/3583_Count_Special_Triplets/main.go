package main

import "slices"

func specialTriplets(nums []int) int {
    // num to sorted indexes in nums array
    num_to_indexes := make(map[int][]int);
    for i, n := range nums {
        if _, found := num_to_indexes[n]; !found {
            num_to_indexes[n] = make([]int, 0);
        }
        num_to_indexes[n] = append(num_to_indexes[n], i);
    }

    const MOD = 1000000007;
    total := 0;

    for k, v := range num_to_indexes {
        if k != 0 {
            pair_array, found := num_to_indexes[k*2];
            // quick out.
            if !found || len(pair_array) <= 1 { continue; }

            for _, n := range v {
                index, ok := slices.BinarySearch(pair_array, n);
                if ok { panic("not possible."); }
                nums_left, nums_right := index, len(pair_array)-index;
                total = (total + nums_left * nums_right) % MOD;
            }
        } else {
            // gotta make a special case for this one.
            //
            // there is proably a formula for this
            for i := 1; i < len(v)-1; i++ {
                left, right := i, len(v) - i - 1;
                total = (total + left * right) % MOD;
            }
        }
    }

    return total;
}