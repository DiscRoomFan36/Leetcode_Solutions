package main

import "slices"

func gcdSum(nums []int) int64 {
    if len(nums) == 0 { panic("what"); }

    prefixed_gcd := make([]int, len(nums));
    maximum := 0;
    for i, n := range nums {
        maximum = max(maximum, n);
        prefixed_gcd[i] = gcd(maximum, n);
    }

    slices.Sort(prefixed_gcd);
    sum := 0;
    for i := range len(prefixed_gcd)/2 {
        j := len(prefixed_gcd)-1-i;
        sum += gcd(prefixed_gcd[i], prefixed_gcd[j]);
    }
    return int64(sum);
}

// https://en.wikipedia.org/wiki/Greatest_common_divisor#Binary_GCD_algorithm
func gcd(a, b int) int {
    if a < 0 || b < 0 { panic("no negative numbers!"); }
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
