package main

import "slices"

func avoidFlood(rains []int) []int {
    answer := make([]int, len(rains));

    const ANSWER_ON_RAINY_DAY = -1;

    // map from pond number to the day when it was filled with water.
    pond_map := make(map[int]int);

    // the index of any drying day
    drying_days := make([]int, 0);

    for i, rain := range rains {
        if rain > 0 {
            // answer is always -1 on rainy days.
            answer[i] = -1;

            // was_in_map aka "was this filled with water before now?"
            day_when_filled_with_water, was_in_map := pond_map[rain];
            if was_in_map {
                // we need to drain it before now.
                // the best day to use would be the first drying day after the pond was filled.
                best_day_to_use_index, found := slices.BinarySearch(drying_days, day_when_filled_with_water);
                if found { panic("not possible."); }

                // check if we find a valid day
                if best_day_to_use_index == len(drying_days) { return []int{}; }

                best_day_to_use := drying_days[best_day_to_use_index]

                // remove the best day. this is a slow operation.
                drying_days = append(drying_days[:best_day_to_use_index], drying_days[best_day_to_use_index+1:]...);

                answer[best_day_to_use] = rain;
            }

            // remember that we just filled a new pond this day.
            pond_map[rain] = i;

        } else {

            // remember when we can dry something
            drying_days = append(drying_days, i);
            // zero is not a valid target to dry.
            answer[i] = 1;
        }
    }

    return answer;
}