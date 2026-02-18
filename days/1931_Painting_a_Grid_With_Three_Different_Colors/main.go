package main

func colorTheGrid(m int, n int) int {
    if n == 0 || m == 0 { return 0; }

    // make 'm' the small one, dose not effect the output.
    if n < m { m, n = n, m; }

    // the trick here is that m is small.
    //
    // at m == 10, there are 3*2^(10-1) == 1536 different wall combos,
    // pretty reasonable amount.
    if m > 10 { panic("m must be small for this to work"); }

    // remember to mod your outputs.
    const MOD = 1000000007;

    type Paint_Color uint;
    const (
        R Paint_Color = iota;
        G;
        B;
    )
    all_colors := [3]Paint_Color{ R, G, B };

    // these are m long.
    type Wall_Section []Paint_Color;

    // 3 * 2^(m-1)
    number_of_wall_combos := 3*(1 << (m-1));
    all_possible_m_wall_combos := make([]Wall_Section, 0, number_of_wall_combos);

    { // recursivly fill 'all_possible_m_wall_combos' with all possible combos.
        var recur func(current_colors []Paint_Color, depth int);
        recur   = func(current_colors []Paint_Color, depth int) {
            if depth == 0 {
                if len(current_colors) != m { panic("must be true"); }

                // we must make a copy here, dont want to get
                // stomped by ourselves in another recursive frame
                new_wall := make([]Paint_Color, m);
                copy(new_wall, current_colors);

                Append(&all_possible_m_wall_combos, Wall_Section(new_wall));
                return;
            }

            // loop over all colors
            for _, color := range all_colors {
                if (len(current_colors) == 0) || (current_colors[len(current_colors)-1] != color) {
                    recur(append(current_colors, color), depth - 1);
                }
            }
        }
        recur([]Paint_Color{}, m);
    }

    if len(all_possible_m_wall_combos) != number_of_wall_combos {
        panic("not mathamatically possible!");
    }

    // calculate witch walls can be put after one another,
    // so we dont have to do this every loop.
    // propbably doesn't matter much, this calculation is very cheap.
    //
    // i personally dont know how to do it, but this problem could
    // be solved in log(n) time, just gotta use the fast
    // matrix power algorithum.
    can_put_wall_after_wall := make_grid[bool](number_of_wall_combos, number_of_wall_combos);
    for i := range number_of_wall_combos {
        this_wall := all_possible_m_wall_combos[i];

        for j := range number_of_wall_combos {
            other_wall := all_possible_m_wall_combos[j];
            
            flag := true;
            for k := range m {
                if this_wall[k] == other_wall[k] {
                    flag = false;
                }
            }
            can_put_wall_after_wall[i][j] = flag;

        }
    }

    curr := make([]uint64, number_of_wall_combos);
    next := make([]uint64, number_of_wall_combos);

    // at n == 1, we have 1 of every wall
    for i := range curr { curr[i] = 1; }

    for range n-1 {
        for i := range number_of_wall_combos {
            next[i] = 0;
            for j := range number_of_wall_combos {
                if can_put_wall_after_wall[i][j] {
                    next[i] += curr[j];
                }
            }
            // were using uint64's, we have alot of room.
            next[i] = next[i] % MOD;
        }

        // stomp curr with next.
        copy(curr, next);
    }

    result := uint64(0);
    for i := range curr { result += curr[i]; }
    return int(result % MOD);
}


func Append[T any](slice *[]T, items ...T) {
	*slice = append(*slice, items...)
}

// makes a grid, with default values. i hope go can turn this into something cool.
//
// might be better to make([]T, m*n) then grab some slices.
func make_grid[T any](n, m int) [][]T {
    grid := make([][]T, m)
    for j := range m { grid[j] = make([]T, n) }
    return grid
}
