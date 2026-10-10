package main

import (
	"cmp"
	"slices"
)

func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
    if len(nums1) == 0 { return 0; }

    for i := range nums1 {
        nums1[i] = Abs(nums1[i] - nums2[i]);
    }
    slices.SortFunc(nums1, func(a, b int) int { return cmp.Compare(-a, -b) });

    // makes the loop ending easyier,
    nums1 = append(nums1, 0);

    answer := 0;

    k := k1 + k2;
    last_number := nums1[0];
    last := 1;
    outer: for ; last < len(nums1); last += 1 {
        if k == 0 {
            answer += Square(last_number) * last;
            break outer;
        }
        for last_number > nums1[last] {
            if k < last {
                answer += Square(last_number-1) * k + Square(last_number) * (last-k);
                break outer;
            }

            last_number -= 1;
            k -= last;
        }
    }

    // grab the rest
    for _, n := range nums1[last:] { answer += Square(n); }
    return int64(answer);
}

func Square(x int) int {
    return x * x;
}

func Abs(x int) int {
    if x < 0 { x = -x; }
    return x;
}
