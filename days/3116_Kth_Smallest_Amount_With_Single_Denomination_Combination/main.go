package main

import "slices"

func findKthSmallest(coins []int, k int) int64 {
    slices.Sort(coins);

    left, right := k, k*coins[0] + 1;

    bit_count := make([]int,   1 << len(coins));
    lcm       := make([]int, 1 << len(coins));

    for mask := 1; mask < (1 << len(coins)); mask += 1 {
        current_lcm := 1;
        for i, coin := range coins {
            if mask >> i & 1 == 0 { continue; }

            this_gcd := gcd(current_lcm, coin);
            temp := current_lcm / this_gcd;
            // if the lcm is bigger than right, we dont care, this can never be the answer.
            if temp <= right / coin {
                current_lcm = temp * coin;
            } else {
                current_lcm = right + 1;
                break;
            }
            bit_count[mask] += 1;
        }
        lcm[mask] = current_lcm;
    }

    count := func(x int) int {
        result := 0;
        for mask := 1; mask < (1 << len(coins)); mask += 1 {
            if lcm[mask] > x { continue; }

            if bit_count[mask] & 1 == 1 {
                result += x / lcm[mask];
            } else {
                result -= x / lcm[mask];
            }
        }
        return result;
    }

    for left < right {
        x := (left + right) / 2;
        if count(x) >= k {
            right = x;
        } else {
            left = x + 1;
        }
    }
    return int64(left);
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
