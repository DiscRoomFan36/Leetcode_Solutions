package main

func numOfWays(n int) int {
    // the base case
    if n <= 0 { return 0; }

    // remember to mod your outputs.
    const MOD = 1000000007;

    // we don't use this in the main loop, just to generate the 'can_put_wall_after_wall'
    type Paint_Color uint;
    const (
        R Paint_Color = iota;
        Y;
        G;
    )

    const WALL_SEGMENTS = 3;
    type Wall_Color [WALL_SEGMENTS]Paint_Color;

    const BASE_CASE_COUNT = 12;
    base_walls := [BASE_CASE_COUNT]Wall_Color{
        { R, Y, R },
        { R, Y, G },
        { R, G, R },
        { R, G, Y },

        { Y, R, Y },
        { Y, R, G },
        { Y, G, R },
        { Y, G, Y },

        { G, R, Y },
        { G, R, G },
        { G, Y, R },
        { G, Y, G },
    };

    // calculate witch walls can be put after one another, so we dont have to do this every loop.
    // propbably doesn't matter much, this calculation is very cheap.
    can_put_wall_after_wall := [BASE_CASE_COUNT][BASE_CASE_COUNT]bool{};
    for i := range BASE_CASE_COUNT {
        this_wall := base_walls[i]

        for j := range BASE_CASE_COUNT {
            other_wall := base_walls[j]

            flag := true;
            for k := range WALL_SEGMENTS {
                if this_wall[k] == other_wall[k] {
                    flag = false;
                    // break; // this loop will be unroll'd
                }
            }
            can_put_wall_after_wall[i][j] = flag;
        }
    }


    // were gonna be doing the swap buffer trick
    curr := [BASE_CASE_COUNT]uint64{}
    next := [BASE_CASE_COUNT]uint64{}

    // base case is all 1's
    for i := range curr { curr[i] = 1; }

    // preform this trick until we do it enough times
    // -1 because were allready on the first step.
    for range n-1 {

        // for every wall in next
        for i := range BASE_CASE_COUNT {
            // reset next.
            next[i] = 0;
            // loop over current and add all that satidy the condition.
            for j := range BASE_CASE_COUNT {
                if can_put_wall_after_wall[i][j] {
                    next[i] += curr[j];
                }
            }
            // only do mod afterwords, 12 * 10^9 cannot overflow 64-bit
            next[i] %= MOD;
        }

        // stomp current.
        curr = next;
    }

    // add all the counts in current, 
    result := uint64(0);
    for i := range curr { result += curr[i]; }
    return int(result % MOD); // only do mod afterwords, 12 * 10^9 cannot overflow 64-bit
}