package main

import (
	"math"
	"slices"
)

func minimumCost(nums []int) int {
    if len(nums) <= 3 { return Sum(nums) }

    // since the first number has to be the start of an array, just get the next smallest 2
    smallest_2 := Smallest_N_Numbers_In_Order(nums[1:], 2)
    return nums[0] + Sum(smallest_2)
}

func Sum(xs []int) int {
    sum := 0
    for _, x := range xs { sum += x }
    return sum
}

// this is O(n*len(nums))
//
// could be made O(log(n)*len(nums)) with a heap.
func Smallest_N_Numbers_In_Order(nums []int, n int) []int {
    if len(nums) < n { panic("array less than n") }

    const INF = math.MaxInt

    // this is a sorted array of the smallest N elements
    smallest_n := make([]int, n)

    // funny
    for i := range smallest_n { smallest_n[i] = nums[i] }
    slices.Sort(smallest_n)

    for j := n; j < len(nums); j++ {
        num := nums[j]

        // check if this number needs to be placed in this array at all.
        if num >= smallest_n[len(smallest_n)-1] { continue }

        // this could be technically faster, if you use a heap, and pop the top when its num is smaller.
        // but i dont like the look of golangs, heap container.
        index_to_insert, _ := slices.BinarySearch(smallest_n, num)
        Fixed_Size_Array_Insert(smallest_n, num, index_to_insert)
    }

    return smallest_n
}

func Fixed_Size_Array_Insert[T any](arr []T, item T, index int) {
    // copy() has memmove() properties, so lets do 1 big copy.
    //
    // if index == len(arr)-1, nothing happens, so its fine. :)
    copy(
        arr[index+1:          ],
        arr[index  :len(arr)-1],
    )

    arr[index] = item
}
