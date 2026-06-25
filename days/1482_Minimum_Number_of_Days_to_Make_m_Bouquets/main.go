package main

func minDays(bloomDay []int, m int, k int) int {
    max_bloom_day := 0;
    for _, day := range bloomDay { max_bloom_day = max(max_bloom_day, day); }

    can_make_bouquets_on_day := func(day int) bool {
        // be greedy
        num_made := 0;
        in_a_row := 0;
        for _, bloom_day := range bloomDay {
            if bloom_day <= day {
                in_a_row += 1;
                if in_a_row >= k {
                    in_a_row = 0;
                    num_made += 1;
                    if num_made >= m { return true; }
                }
            } else {
                in_a_row = 0;
            }
        }
        return false;
    }

    // if its not possible on the last day, its never possible.
    //
    // could use math to figure this out faster.
    if !can_make_bouquets_on_day(max_bloom_day) { return -1; }

    // binary search
    low, high := 1, max_bloom_day;
    for low < high {
        mid := (low + high) / 2;
        if can_make_bouquets_on_day(mid) {
            high = mid;
        } else {
            low = mid + 1;
        }
    }
    return high;
}
