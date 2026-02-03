package main

func isTrionic(nums []int) bool {
    j := changes_for_how_long(nums[   :], true )
    if j == 0 { return false }

    k := changes_for_how_long(nums[j  :], false)
    if k == 0 { return false }

    l := changes_for_how_long(nums[j+k:], true )
    if l == 0 { return false }

    // check if we reached the end of the array.
    n := len(nums)
    if j+k+l == n-1 { return true }

    return false
}

func changes_for_how_long(nums []int, increasing bool) int {
    for i := range len(nums)-1 {
        next_decreasing := nums[i] > nums[i+1]
        next_increasing := nums[i] < nums[i+1]

        doing_what_it_should := (increasing && next_increasing) || (!increasing && next_decreasing)

        if !doing_what_it_should { return i }
    }
    return len(nums)-1
}
