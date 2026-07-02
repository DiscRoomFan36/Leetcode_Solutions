package main

func findSafeWalk(grid [][]int, health int) bool {
    m, n := len(grid), len(grid[0]);

    // this is a weird state, just handle it here.
    if grid[0][0] == 1 { health -= 1; }
    if health == 0 { return false; }

    type Vec2 struct { x, y int; };

    health_to_paths := make([][]Vec2, health);
    pos_to_best_health_on_tile := make_grid[int](n, m);

    // pos_to_best_health_on_tile[0][0] = health;
    // inital point.
    health_to_paths[health-1] = []Vec2{ {0, 0} };

    for current_health := health; current_health > 0; current_health-- {
        paths := &health_to_paths[current_health-1];

        for len(*paths) > 0 {
            pos := Pop(paths);

            // we reached the end. cold path.
            if pos.x == n-1 && pos.y == m-1 { return true; }

            if pos_to_best_health_on_tile[pos.y][pos.x] >= current_health { continue; }
            pos_to_best_health_on_tile[pos.y][pos.x] = current_health;

            // go to all other tiles.
            add_other_pos := func(pos Vec2, x, y int) {
                pos.x += x;
                pos.y += y;
                // bounds checking
                if pos.x < 0 || n <= pos.x { return; }
                if pos.y < 0 || m <= pos.y { return; }

                this_health := current_health;
                // take 1 damage
                if grid[pos.y][pos.x] == 1 { this_health -= 1; }

                // check if you died.
                if this_health == 0 { return; }

                Append(&health_to_paths[this_health-1], pos);
            }

            add_other_pos(pos,  1,  0);
            add_other_pos(pos, -1,  0);
            add_other_pos(pos,  0,  1);
            add_other_pos(pos,  0, -1);
        }
    }

    return false;
}


func Append[T any](slice *[]T, items ...T) *T {
    *slice = append(*slice, items...);
    return &(*slice)[len(*slice)-1];
}

func Pop[T any](slice *[]T) T {
    item  := (*slice)[ len(*slice)-1];
    *slice = (*slice)[:len(*slice)-1];
    return item;
}

// makes a grid, with default values. i hope go can turn this into something cool.
//
// might be better to make([]T, m*n) then grab some slices.
func make_grid[T any](n, m int) [][]T {
    grid := make([][]T, m);
    for j := range m { grid[j] = make([]T, n); }
    return grid;
}
