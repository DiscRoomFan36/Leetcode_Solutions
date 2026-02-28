package main

import "math/bits"

func concatenatedBinary(_n int) int {
    if _n <= 0 { panic("this is a leetcode problem."); }
    n := uint(_n);

    const MOD = 1000000007;

    number := uint(0);
    for i := uint(1); i <= n; i++ {
        // we could do a funny calculation here, to skip this
        // function call, and just inc the number when its a
        // power of 2, however i dont know what this funciton
        // compiles down to, it might be a single machine
        // instruction. who knows. so here it stays,
        //
        // just found out this is just a regular function.
        // would totaly be faster to keep track.
        length_of_new_number := bits.Len(i);

        // you know, i dont fully internalize the reason this
        // MOD works, i guess mutiplying by 2^len is fine?
        number = (number << length_of_new_number | i) % MOD;
    }
    return int(number);
}
