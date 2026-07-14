package main

func minimumSeconds(nums []int) int {
    num_to_indexes := make(map[int][]int);
    for i, n := range nums {
        num_to_indexes[n] = append(num_to_indexes[n], i);
    }

    // the base case (if best_count == 1,)
    longest_distance_from_num := len(nums)/2;
    for _, indexes := range num_to_indexes {
        if len(indexes) == 1 { continue; }

        longest_distance_from_this_num := 0;
        for i := 1; i < len(indexes); i++ {
            last_index := indexes[i-1];
            index      := indexes[i  ];

            dist := (index - last_index) - 1;
            biggest_score := (dist + 1) / 2;
            longest_distance_from_this_num = max(longest_distance_from_this_num, biggest_score);
        }

        // handle the wrap around.
        longest_distance_from_this_num = max(
            longest_distance_from_this_num,
            ((indexes[0] + len(nums) - 1) - indexes[len(indexes)-1] + 1) / 2,
        );

        longest_distance_from_num = min(longest_distance_from_num, longest_distance_from_this_num);
    }

    return longest_distance_from_num;
}
