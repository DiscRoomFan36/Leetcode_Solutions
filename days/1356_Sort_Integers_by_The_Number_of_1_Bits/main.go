package main

import (
	"cmp"
	"math/bits"
	"slices"
)

func sortByBits(arr []int) []int {
    slices.SortFunc(arr,
        func(a, b int) int {
            a_bits := bits.OnesCount(uint(a));
            b_bits := bits.OnesCount(uint(b));
            if a_bits != b_bits {
                return cmp.Compare(a_bits, b_bits);
            } else {
                return cmp.Compare(a, b);
            }
        },
    );
    return arr;
}