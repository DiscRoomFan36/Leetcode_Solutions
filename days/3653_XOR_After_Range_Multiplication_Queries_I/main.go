package main

func xorAfterQueries(nums []int, queries [][]int) int {
    const MOD = 1000000007;

    for _, q := range queries {
        l, r, k, v := q[0], q[1], q[2], q[3];

        for i := l; i <= r; i += k {
            nums[i] = (nums[i] * v) % MOD;
        }
    }

    result := 0;
    for _, n := range nums {
        result ^= n;
    }
    return result;
}
