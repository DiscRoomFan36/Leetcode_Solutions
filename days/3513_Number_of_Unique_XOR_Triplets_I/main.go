package main

import "math/bits"

func uniqueXorTriplets(nums []int) int {
    n := len(nums);

    // because we know nums is a permutation from 1..n,
    //
    // and the xor of a number with itself is 0,
    //
    // its easy to get all numbers <= n.

    if n <= 2 { return n;   }
    // we can make zero when n >= 3.
    if n <= 3 { return n+1; }

    // now we need to handle the case where we can get numbers bigger than n...
    //
    // is it possible to get all numbers upto (n*2)-1?

    // the first number has to have the leftmost bit set, but after that, 2-3 different numbers 

    // 100000
    // 011111
    // 0xxxxx
    //
    // 100000
    // 011110
    // 000001

    // we can make every number upto n*2 - 1
    pow2 := 64 - bits.LeadingZeros64(uint64(n));
    return (1 << pow2) - 1 + 1; // +1 accounts for zero.
}
