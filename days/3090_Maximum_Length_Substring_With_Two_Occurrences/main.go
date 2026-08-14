package main

func maximumLengthSubstring(s string) int {
    return maxSubarrayLength(s, 2);
}

func maxSubarrayLength(nums string, k int) int {
    best_length := min(k, len(nums));
    if len(nums) == 0 { return best_length; }

    // index's into the nums array, defines a subarray.
    head, tail := 0, 0;

    // keep track of elements encountered.
    element_counts := make(map[byte]int);
    // add first number, assumes k > 0.
    element_counts[nums[0]] += 1;

    for {
        // shift tail forwards
        tail += 1;
        // end condition
        if tail >= len(nums) { break; }
        // add new element to counts
        element_counts[nums[tail]] += 1;
        // move head forwards until array is valid again
        for element_counts[nums[tail]] > k {
            head_element := nums[head];
            element_counts[head_element] -= 1;
            head += 1;
        }

        best_length = max(best_length, tail - head + 1);
    }

    return best_length;
}
