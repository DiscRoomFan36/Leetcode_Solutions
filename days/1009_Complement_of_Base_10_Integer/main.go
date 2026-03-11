package main

func bitwiseComplement(n int) int {
    // special case this.
    if n == 0 { return 1; }

    next_pow_2 := next_pow_of_2(n);
    return ^n & (next_pow_2-1)
}

func next_pow_of_2(v int) int {
    v--;
    v |= v >>  1;
    v |= v >>  2;
    v |= v >>  4;
    v |= v >>  8;
    v |= v >> 16;
    v |= v >> 32;
    v++;
    return v;
}