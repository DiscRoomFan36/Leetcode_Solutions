package main

import "math"

func swimInWater(grid [][]int) int {    
    if len(grid) == 0 { return 0 }
    n := len(grid);

    type Pos struct { x, y int }

    have_been_to := make_grid[bool](n, n);

    // must be at minimum the time it would take to get to the end, or away from the begining.
    water_level := max(grid[0][0], grid[n-1][n-1]);

    positions_to_check := []Pos{ {0, 0} };
    positions_to_check_next_time := []Pos{};

    for {
        // fmt.Println("water_level:", water_level)
        const INF = math.MaxInt
        min_cannot_reach := INF

        for len(positions_to_check) > 0 {
            pos := Pop(&positions_to_check);

            // might be fast to have a second check, something like "seen this before"
            // so that things that are entered into next_time dont get flooded.
            // 
            // on the other hand, entries can only get entered into this list at most 4 times,
            // one for each edge that surrounds the cell.
            if have_been_to[pos.y][pos.x] { continue; }

            ground_height := grid[pos.y][pos.x];

            // if the level of ground is to high to swim on, check it later
            if ground_height > water_level {
                min_cannot_reach = min(min_cannot_reach, ground_height)
                Append(&positions_to_check_next_time, pos);
                continue;
            }

            have_been_to[pos.y][pos.x] = true;

            // there is probably a better place to put this check,
            // but here is the most convienent place.
            if pos.x == n-1 && pos.y == n-1 { return water_level; }

            // make sure we dont go out of bounds, and flood fill.
            if pos.x > 0   { Append(&positions_to_check, Pos{ pos.x-1, pos.y   }); }
            if pos.x < n-1 { Append(&positions_to_check, Pos{ pos.x+1, pos.y   }); }
            if pos.y > 0   { Append(&positions_to_check, Pos{ pos.x  , pos.y-1 }); }
            if pos.y < n-1 { Append(&positions_to_check, Pos{ pos.x  , pos.y+1 }); }
        }

        if len(positions_to_check_next_time) == 0 { panic("not possible"); }
        // swap buffers, now positions_to_check_next_time is empty.
        positions_to_check, positions_to_check_next_time = positions_to_check_next_time, positions_to_check

        // we didnt find the end. increase the water level
        if min_cannot_reach == INF { panic("not possible") }
        water_level = min_cannot_reach
    }
}


// makes a grid, with default values. i hope go can turn this into something cool.
//
// might be better to make([]T, m*n) then grab some slices.
func make_grid[T any](n, m int) [][]T {
    grid := make([][]T, m)
    for j := range m { grid[j] = make([]T, n) }
    return grid
}

func Pop[T any](slice *[]T) T {
	item  := (*slice)[ len(*slice)-1]
	*slice = (*slice)[:len(*slice)-1]
	return item
}

func Append[T any](slice *[]T, items ...T) {
	*slice = append(*slice, items...)
}
