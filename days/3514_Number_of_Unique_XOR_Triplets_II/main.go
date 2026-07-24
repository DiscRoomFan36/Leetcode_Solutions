package main

import "math/bits"

func uniqueXorTriplets(nums []int) int {
    max_num := 0;
    for _, n := range nums { max_num = max(max_num, n); }

    max_posible_xor := 1 << (64 - bits.LeadingZeros64(uint64(max_num)));

    set  := make([]bool, max_posible_xor);
    // zero is always possible, by xor'ing with self.
    set[0] = true;

    for i := 0; i < len(nums); i++ {
        for j := i+1; j < len(nums); j++ {
            set[nums[i] ^ nums[j]] = true;
        }
    }

    final := make([]bool, max_posible_xor);
    for _, n := range nums {
        for i, b := range set {
            if b {
                final[n ^ i] = true;
            }
        }
    }

    final_count := 0;
    for _, b := range final {
        if b { final_count += 1; }
    }

    return final_count;
}
