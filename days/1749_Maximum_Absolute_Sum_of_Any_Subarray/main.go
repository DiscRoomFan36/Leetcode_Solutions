package main

func maxAbsoluteSum(nums []int) int {
    res_max := 0
    res_min := 0

    max_ending := 0
    min_ending := 0

    for _, n := range nums {
        // Either extend the previous subarray or start 
        // new from current element
        if max_ending + n > n { max_ending += n
        } else                { max_ending  = n }
        if min_ending + n < n { min_ending += n
        } else                { min_ending  = n }

        // Update result if the new subarray sum is larger
        res_max = max(res_max, max_ending)
        res_min = min(res_min, min_ending)
    }

    return max(res_max, -res_min)
}