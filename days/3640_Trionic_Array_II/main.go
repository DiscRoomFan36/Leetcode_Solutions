package main

func maxSumTrionic(nums []int) int64 {
    n := len(nums)

    const INF = 1 << 60

    // this is all the info we need, just need to gather the data, then process it.
    type Tronic_Part struct {
        // you might think we can get away with only one
        // of these numbers, and just checking the before
        // and after, this is not the case. because we still
        // have to deal with duplicate elements. I suppose
        // it could be encoded some way into the array, but meh.
        start_index, end_index int
    }

    parts := make([]Tronic_Part, 0)

    // flip flops between checking for increaseing
    // numbers, and decreasing numbers.
    flip_flop := true
    for index := 0; index < n-1; {
        length := changes_for_how_long(nums[index:], flip_flop)

        parts = append(parts, Tronic_Part{
            start_index : index,
            end_index   : index + length,
        })

        index     += length
        flip_flop  = !flip_flop

        // if 2 elements are equal, this skips over it.
        if (length == 0) && (index < n-1) && (nums[index] == nums[index+1]) {
            // skip over all duplicate numbers, remember to leave index at the last duplicate.
            duplicate_number := nums[index]
            index += 1
            for (index < n-1) && (nums[index+1] == duplicate_number) {
                index += 1
            }
        }
    }

    max_tronic := -INF
    // iterate over every second one. we know every even one is an accending one.
    for i := 0; i < len(parts)-2; i += 2 {
        // make sure no invalid arrays come this way.
        // invalid slices have length == 0
        if parts[i  ].end_index - parts[i  ].start_index == 0 { continue }
        if parts[i+1].end_index - parts[i+1].start_index == 0 { continue }
        if parts[i+2].end_index - parts[i+2].start_index == 0 { continue }

        total_score := 0

        { // get the sum of the middle segment, gonna have to do this anyway. why not first?
            start_index := parts[i+1].start_index
            end_index   := parts[i+1].end_index

            sum := 0
            for _, n := range nums[start_index:end_index] { sum += n }
            total_score += sum
        }

        // it might be better to only take parts of the acending parts
        // from the tronic array, so lets find out how
        // much we can lose. or what is the best amount to take.

        { // 1st ascending, in reverse.
            start_index := parts[i].start_index
            end_index   := parts[i].end_index

            // make sure we dont count the end one, thats taken care of in the middle sum
            running_score := nums[end_index-1]
            best_score    := running_score

            for j := end_index-2; j >= start_index; j-- {
                running_score += nums[j]
                best_score = max(best_score, running_score)
            }

            total_score += best_score
        }

        { // 2nd ascending, forwards
            start_index := parts[i+2].start_index
            end_index   := parts[i+2].end_index

            // remember to count the start one, also needs to be at least 2 long, so its real.
            running_score := nums[start_index] + nums[start_index+1]
            best_score    := running_score

            // (j <= end_index) to count end as well
            for j := start_index+2; j <= end_index; j++ {
                running_score += nums[j]
                best_score = max(best_score, running_score)
            }

            total_score += best_score
        }

        if total_score > max_tronic {
            max_tronic = total_score
        }
    }
    return int64(max_tronic)
}

func changes_for_how_long(nums []int, increasing bool) int {
    for i := range len(nums)-1 {
        next_decreasing := nums[i] > nums[i+1]
        next_increasing := nums[i] < nums[i+1]

        doing_what_it_should := (increasing && next_increasing) || (!increasing && next_decreasing)

        if !doing_what_it_should { return i }
    }
    return len(nums)-1
}
