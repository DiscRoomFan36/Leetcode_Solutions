package main

func minCost(colors string, neededTime []int) int {
    n := len(colors);
    
    total_cost := 0;

    // find the runs of the same color

    start_of_run := 0;
    for start_of_run < n {
        start_color := colors[start_of_run];

        end_of_run := start_of_run;

        // keep a running sum and max for speed, also binding energy.
        sum, max := neededTime[start_of_run], neededTime[start_of_run];
        for i := end_of_run + 1; i < n; i++ {
            if colors[i] != start_color { break; }

            time := neededTime[i]
            sum += time;
            if max < time { max = time; }

            end_of_run = i;
        }

        // we remove everything that isnt the biggest one, if there is only one, we keep it.
        total_cost += sum - max;

        start_of_run = end_of_run + 1;
    }

    return total_cost;
}