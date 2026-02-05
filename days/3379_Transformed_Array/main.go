package main

func constructTransformedArray(nums []int) []int {
    n := len(nums)
    result := make([]int, n)

    for i := range n {
        index_altered := i + nums[i] 
        // proper mod
        index_mod_n := ((index_altered % n) + n) % n
        result[i] = nums[index_mod_n]
    }

    return result
}