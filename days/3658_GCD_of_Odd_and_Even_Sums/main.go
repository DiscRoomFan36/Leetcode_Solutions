package main

func gcdOfOddEvenSums(n int) int {
    sum_odd  := 0;
    sum_even := 0;
    for i := range n {
        sum_odd  += i*2 + 1;
        sum_even += i*2 + 2;
    }

    return int(gcd(uint(sum_odd), uint(sum_even)));
}

// https://en.wikipedia.org/wiki/Greatest_common_divisor#Binary_GCD_algorithm
func gcd(a, b uint) uint {
    if a == 0 { return b; }
    if b == 0 { return a; }

    // find shared powers of 2
    d := 0;
    for (a % 2 == 0) && (b % 2 == 0) {
        a /= 2;
        b /= 2;
        d += 1;
    }

    // make both odd
    for a % 2 == 0 { a /= 2; }
    for b % 2 == 0 { b /= 2; }

    // until a == b
    for a != b {
        if a < b { a, b = b, a; }

        a = (a - b) / 2;
        for a % 2 == 0 { a /= 2; }
    }

    return (1 << d) * a;
}
