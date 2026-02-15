package main

func countTriples(_n int) int {
    if _n <= 0 { panic("why") }
    n := uint(_n)

    result := 0
    for a := uint(1); a <= n; a++ {
        for b := a+1; b <= n; b++ {
            c_sqr := a*a + b*b
            // its to big
            if c_sqr > n*n { break }
            // binary search for speed here.
            maybe_c := int_sqrt(c_sqr)
            if maybe_c*maybe_c == c_sqr {
                result += 2
            }
        }
    }
    return result
}

// https://en.wikipedia.org/wiki/Integer_square_root
func int_sqrt(n uint) uint {
    L, R := uint(1), n
    for (L < R) {
        R = L + ((R - L) / 2);
        L = n / R;
    }
    return R;
}