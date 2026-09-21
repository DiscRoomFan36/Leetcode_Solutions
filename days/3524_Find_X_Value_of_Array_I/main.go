package main

func resultArray(nums []int, k int) []int64 {
    result := make([]int64, k);

    curr := make([]int64, k);
    prev := make([]int64, k);
    for _, n := range nums {
        // this is the subarray consisting of only this element
        curr[n % k] += 1;

        // continue the older arrays,
        for j := range k {
            r := (j * n) % k;
            curr[r] += prev[j];
        }

        for j := range k {
            // accumulate for the result;
            result[j] += curr[j];
            // swap buffers and clear curr buffer
            prev[j] = curr[j];
            curr[j] = 0;
        }
    }

    return result;
}
