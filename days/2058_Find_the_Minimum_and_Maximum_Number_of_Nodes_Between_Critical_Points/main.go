package main

import "math"

type ListNode struct {
    Val int
    Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
    // i know this problem is trying to get me to do it in one go,
    // but i dont actually care.
    //
    // also a temperary allocator will rip thought this entire thing.

    nums := make([]int, 0);
    for node := head; node != nil; node = node.Next {
        nums = append(nums, node.Val);
    }

    critical_points := make([]int, 0);
    for i := 1; i < len(nums)-1; i += 1 {
        is_critical_point := false;
        if nums[i-1] < nums[i] && nums[i] > nums[i+1] { is_critical_point = true; }
        if nums[i-1] > nums[i] && nums[i] < nums[i+1] { is_critical_point = true; }

        if is_critical_point {
            critical_points = append(critical_points, i);
        }
    }

    if len(critical_points) < 2 { return []int{-1, -1}; }

    min_distance := math.MaxInt;
    for i := 1; i < len(critical_points); i += 1 {
        dist := critical_points[i] - critical_points[i-1];
        min_distance = min(min_distance, dist);
    }
    return []int{min_distance, critical_points[len(critical_points)-1] - critical_points[0]};
}
