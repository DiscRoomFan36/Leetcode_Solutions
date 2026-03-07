package main

import "math/bits"

func minBitwiseArray(nums []int) []int {
    result := make([]int, len(nums))

    for_all_numbers: for i := range nums {
        n := nums[i]

        // when a number is even, it is not possible to make it with   X | (X+1)   as this action always sets the first bit.
        if n % 2 == 0 {
            result[i] = -1
            continue for_all_numbers;
        }

        // (n XOR 1) is a valid soltion, but not the smallest 1
        //
        // 0b000010111 lets say this is a prime (in binary)
        //
        // 0b000010110 is a valid solution
        //
        // 0b000010011 this is also valid, and smaller
        //

        // find the first zero bit from the left.
        first_zero_index := bits.TrailingZeros(^uint(n));
        // set the bit to the right to zero.
        result[i] = n & ^(1 << (first_zero_index-1));
    }

    return result
}
