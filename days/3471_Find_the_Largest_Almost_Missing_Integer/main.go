package main

func largestInteger(nums []int, k int) int {

    if k == len(nums) {
        best_num := -1;
        for _, num := range nums {
            best_num = max(best_num, num);
        }
        return best_num;
    }

    num_seen_count := make(map[int]int);
    for start := range len(nums)-k+1 {
        for i := range k {
            num := nums[start + i];
            num_seen_count[num] += 1;
        }
    }

    best_num := -1;
    for num, count := range num_seen_count {
        if count == 1 {
            best_num = max(best_num, num);
        }
    }
    return best_num;
}
