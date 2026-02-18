package main

func hasAlternatingBits(n int) bool {
    // I mean... yes?
    if n == 0 { return true; }

    // shift by 1, then xor, if the bits are alternating, x should be all 1's
    // if not alternating, there will be some gaps, (aka 0's) in the number.
    x := n ^ (n >> 1);
    // y should now be a mutiple of 2, (if it was alternating)
    // or in other words, only have 1 bit set.
    y := x + 1;
    // this removes the lowest bit from the number 
    z := y & (y-1);
    // if z == 0, it was a power of 2, it had alternating bits.
    return z == 0;
}