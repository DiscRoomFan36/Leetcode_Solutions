package main

func stoneGameV(stoneValue []int) int {
    prefix_array := make([]int, len(stoneValue)+1);
    for i := range len(stoneValue) {
        prefix_array[i+1] = prefix_array[i] + stoneValue[i];
    }

    type Range struct { left, right int; };
    range_to_best_score := make(map[Range]int);

    sum_of_range := func(this_range Range) int {
        return prefix_array[this_range.right] - prefix_array[this_range.left];
    }

    var recur func(this_range Range) int;
    recur = func(this_range Range) int {
        // empty range means game over;
        if this_range.left == this_range.right { return 0; }

        score, found := range_to_best_score[this_range];
        if found { return score; }

        // calculate best split
        best_score := 0;
        // (+ 1) & (! <=) is because taking a zero length range is stupid, and ends the game.
        for i := this_range.left + 1; i < this_range.right; i += 1 {
            left_range  := Range{ this_range.left, i  };
            right_range := Range{ i, this_range.right };

            left_sum  := sum_of_range(left_range );
            right_sum := sum_of_range(right_range);

            if left_sum  <= right_sum {
                // its kinda like alpha-beta pruneing.
                if left_sum >= best_score / 2 {
                    best_score = max(best_score, left_sum  + recur(left_range) );
                }
            }
            if right_sum <= left_sum  {
                if right_sum >= best_score / 2 {
                    best_score = max(best_score, right_sum + recur(right_range));
                }
            }
        } 

        range_to_best_score[this_range] = best_score;
        return best_score
    }

    return recur(Range{0, len(stoneValue)});
}
