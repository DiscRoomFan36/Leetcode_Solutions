package main

import (
	"fmt"
	"math/bits"
)

func readBinaryWatch(turnedOn int) []string {
    result := []string{}

    // lets just enumerate all the possibilites.
    // there are only 1024 of them
    //
    // there is probably some better way to iterate the permutations,
    // i should make a generator function to do that.
    for i := range 1 << 10 {
        if bits.OnesCount(uint(i)) != turnedOn { continue }

        bottom_6 :=  i & 0b0000111111
        top_4    := (i & 0b1111000000) >> 6

        // check if the number it would make is invalid.
        if (top_4 > 11) || (bottom_6 > 59) { continue }

        result = append(result, fmt.Sprintf("%d:%02d", top_4, bottom_6))
    }

    return result
}
