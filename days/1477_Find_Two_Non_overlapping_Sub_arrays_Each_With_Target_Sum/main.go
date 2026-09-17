package main

import "math"

func minSumOfLengths(arr []int, target int) int {
    n := len(arr);

    const INF = math.MaxInt / 16;

    smallest_subarrays := make([]int, n+1);
    for i := range smallest_subarrays { smallest_subarrays[i] = INF; }

    answer := INF;

    // useing a sliding window to make small subarrays that sum to target, 
    left := 0;
    running_sum := 0;
    for right := 0; right < n; right += 1 {
        running_sum += arr[right];
        for running_sum > target {
            running_sum -= arr[left];
            left += 1;
        }
        // carry forwards our smallest subarray
        smallest_subarrays[right+1] = smallest_subarrays[right];

        if running_sum == target {
            length := right - left + 1;
            smallest_previously := smallest_subarrays[left];
            if smallest_previously != INF {
                answer = min(answer, length + smallest_previously);
            }
            smallest_subarrays[right+1] = min(smallest_subarrays[right+1], length);
        }
    }

    if answer == INF { return -1; }
    return answer;
}