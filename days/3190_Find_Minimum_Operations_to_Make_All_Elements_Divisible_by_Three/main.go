package main

func minimumOperations(nums []int) int {
    result := 0
    for _, n := range nums {
        // all numbers only need one operation to get to three.
        if n % 3 != 0 { result += 1 }
    }
    return result
}