package main

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
    type Enum uint8;
    const (
        UNGUARDED Enum = iota;
        GUARDED;
        BLOCKER;
    );

    grid := make_grid[Enum](n, m);

    for i := range walls {
        row, col := walls[i][0], walls[i][1];
        grid[row][col] = BLOCKER;
    }

    for i := range guards {
        row, col := guards[i][0], guards[i][1];
        grid[row][col] = BLOCKER;
    }

    for i := range guards {
        row, col := guards[i][0], guards[i][1];

        // sets all cells to GUARDED if not blocked by something.
        mark_until_hits_a_wall := func(start_x, start_y, inc_x, inc_y int) {
            x, y := start_x, start_y;
            for 0 <= x && x < n && 0 <= y && y < m {
                switch grid[y][x] {
                    case UNGUARDED: { grid[y][x] = GUARDED; }
                    case   GUARDED: { /* grid[y][x] = GUARDED; */ }
                    case BLOCKER: { return; }
                }
                x += inc_x;
                y += inc_y;
            }
        }

        mark_until_hits_a_wall(col+1, row  ,  1,  0);
        mark_until_hits_a_wall(col-1, row  , -1,  0);
        mark_until_hits_a_wall(col  , row+1,  0,  1);
        mark_until_hits_a_wall(col  , row-1,  0, -1);
    }

    total_unguarded := 0;
    for j := range m {
        for i := range n {
            if grid[j][i] == UNGUARDED {
                total_unguarded += 1;
            }
        }
    }
    return total_unguarded;
}

// makes a grid, with default values. i hope go can turn this into something cool.
//
// might be better to make([]T, m*n) then grab some slices.
func make_grid[T any](n, m int) [][]T {
    grid := make([][]T, m);
    for j := range m { grid[j] = make([]T, n); }
    return grid;
}
