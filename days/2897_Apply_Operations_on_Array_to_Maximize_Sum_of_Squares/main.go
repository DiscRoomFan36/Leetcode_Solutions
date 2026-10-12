package main

func maxSum(nums []int, k int) int {
    const MOD = 1000000007;

    bit_counts := [64]int{};
    for _, n := range nums {
        for bit := range bit_counts {
            bit_counts[bit] += (n >> bit) & 1;
        }
    }

    answer := 0;
    for range k {
        number := 0;
        for bit, count := range bit_counts {
            if count > 0 {
                number += 1 << bit;
                bit_counts[bit] -= 1;
            }
        }
    
        answer = (answer + number * number) % MOD;
    }
    
    return answer;
}
