package main

import "math/bits"

func countPrimeSetBits(left int, right int) int {
    // numbers less than 10^6 can have at max 19 bits set.
    //
    // 8 primes, what a nice number...
    primes := [...]int{2, 3, 5, 7, 11, 13, 17, 19};

    result := 0;
    for i := left; i <= right; i++ {
        set_bits := bits.OnesCount(uint(i));

        // this will be super fast, this entire thing is gonna be unrolled probably.
        is_prime := false;
        for _, prime := range primes {
            if set_bits == prime {
                is_prime = true;
                break;
            }
        }
        if is_prime { result += 1; }
    }

    return result;
}
