package main

import "math/bits"

func isPowerOfFour(n int) bool {
    if n <= 0 { return false; }

    num_bits := bits.OnesCount(uint(n));
    // there must only be 1 bit in a power of 4.
    if num_bits != 1 { return false; }

    trailing_zeros := bits.TrailingZeros(uint(n));
    if trailing_zeros % 2 != 0 { return false; }

    return true;
}
