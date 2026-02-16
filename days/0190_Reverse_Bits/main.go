package main

func reverseBits(n int) int {
    const (
        a = 0xFFFF0000
        b = 0xFF00FF00
        c = 0xF0F0F0F0
        d = 0xCCCCCCCC
        e = 0xAAAAAAAA
    )

    // hackers delight moment.
    n = ((n & a) >> 16) | ((n << 16) & a)
    n = ((n & b) >>  8) | ((n <<  8) & b)
    n = ((n & c) >>  4) | ((n <<  4) & c)
    n = ((n & d) >>  2) | ((n <<  2) & d)
    n = ((n & e) >>  1) | ((n <<  1) & e)

    return n
}