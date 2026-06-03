package main

import "math"

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
    get_best_end_time := func(start_times, durations []int) int {
        best_end_time := math.MaxInt;
        for i := range len(start_times) {
            best_end_time = min(best_end_time, start_times[i] + durations[i]);
        }
        return best_end_time;
    }

    best_land_end_time  := get_best_end_time( landStartTime,  landDuration);
    best_water_end_time := get_best_end_time(waterStartTime, waterDuration);

    get_best_final_time := func(current_time int, start_times, durations []int) int {
        best_final_time := math.MaxInt;
        for i := range len(start_times) {
            finish_time := max(current_time, start_times[i]) + durations[i];
            best_final_time = min(best_final_time, finish_time);
        }
        return best_final_time;
    }

    return min(
        get_best_final_time(best_land_end_time,  waterStartTime, waterDuration),
        get_best_final_time(best_water_end_time,  landStartTime,  landDuration),
    );
}
