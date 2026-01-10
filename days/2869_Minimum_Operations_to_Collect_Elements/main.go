
package main

func minOperations(nums []int, k int) int {
    seen := make([]bool, k)
    unseen := k

    for i := len(nums)-1; i >= 0; i-- {
        n := nums[i]
        if n > k { continue }
        // if n == 0 { panic("this is a leetcode poblem") }
        if !seen[n-1] {
            seen[n-1] = true
            unseen -= 1
            if unseen == 0 { return len(nums) - i }
        }
    }

    panic("this is a leetcode problem, this cannot happen")
}
