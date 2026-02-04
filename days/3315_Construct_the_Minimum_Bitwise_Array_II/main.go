package main

func minBitwiseArray(nums []int) []int {
    result := make([]int, len(nums))

    for_all_numbers: for i := range nums {
        n := nums[i]

        // when a number is even, it is not possible to make it with   X | (X+1)   as this action allways sets the first bit.
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

        // find the first bit that is not zero.
        for j := 1; j < 64; j++ {
            if n & (1 << j) == 0 {
                // this means the last bit was last.
                // toggle that one bit, that is the smallest possible answer
                result[i] = n ^ (1 << (j-1))
                continue for_all_numbers;
            }
        }
        // this means we git a number with all 1's, for this leetcode problem, we do not worry about that.
        panic("not possible, 0xFFFFFFFF is not a prime")
    }

    return result
}
