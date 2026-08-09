package main

func limitOccurrences(nums []int, k int) []int {
    result := make([]int, 0);
    last_seen := 0;
    last_seen_count := 0;
    for _, n := range nums {
        if last_seen == n {
            last_seen_count += 1;
        } else {
            last_seen = n;
            last_seen_count = 1;
        }

        if last_seen_count <= k {
            result = append(result, n);
        }
    }
    return result;
}
