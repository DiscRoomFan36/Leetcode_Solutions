package main

func canMakeEqual(nums []int, k int) bool {
    can_make_it_equal_to := func(trying_to_equal int, k int) bool {
        next_is_flipped := false;
        for i := range len(nums)-1 {
            num := nums[i];
            if next_is_flipped { num *= -1; }
            next_is_flipped = false;

            if num != trying_to_equal {
                // ran out of flips.
                if k == 0 { return false; }
                k -= 1;
                next_is_flipped = true;
            }
        }

        // check last num. we cannot flip it seperatly.
        final_num := nums[len(nums)-1];
        if next_is_flipped { final_num *= -1; }
        if final_num != trying_to_equal { return false; }

        return true; 
    }

    if can_make_it_equal_to( 1, k) { return true; }
    if can_make_it_equal_to(-1, k) { return true; }
    return false;
}
