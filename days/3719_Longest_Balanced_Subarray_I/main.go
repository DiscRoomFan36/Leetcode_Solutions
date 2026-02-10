package main

func longestBalanced(nums []int) int {
    n := len(nums)

    max_subarray := 0
    for i := range n {
        if max_subarray > n - i { break }

        digits   := make(map[int]int)
        num_odd  := 0
        num_even := 0

        add_number := func(num int) {
            count := digits[num]
            if count == 0 {
                if num % 2 == 0 { num_even += 1
                } else          { num_odd  += 1 }
            }
            digits[num] += 1
        }

        for j := range n-i {
            // TODO we can break early when there is not enough room to make it balanced.
            add_number(nums[i+j])
            if num_odd == num_even { max_subarray = max(max_subarray, j+1) }
        }
    }

    return max_subarray
}