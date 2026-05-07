package main

import "math/bits"

func isFascinating(n int) bool {
    seen_bitfield := uint(0);

    do_test := func(n int) bool {
        // number have exactly 3 digits.
        for range 3 {
            digit := n % 10;
            if digit == 0 { return false; }
            seen_bitfield |= 1 << digit;

            n /= 10;
        }
        return true;
    }

    if !do_test(n  ) { return false; }
    if !do_test(n*2) { return false; }
    if !do_test(n*3) { return false; }

    return bits.OnesCount(seen_bitfield) == 9;
}
