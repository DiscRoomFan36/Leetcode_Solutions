package main

func hasValidPath(grid [][]int) bool {
    if len(grid) == 0 { return false; }
    m, n := len(grid), len(grid[0]);

    // the positions weve been too, to help find when we are in a cycle.
    been_to := make_grid[bool](n, m);

    type Direction int;
    const (
        UP    Direction = 1 << 0;
        LEFT  Direction = 1 << 1;
        DOWN  Direction = 1 << 2;
        RIGHT Direction = 1 << 3;
    );

    path_connections := [6]Direction{
             LEFT |        RIGHT,
        UP |        DOWN,
             LEFT | DOWN,
                    DOWN | RIGHT,
        UP | LEFT,
        UP |               RIGHT,
    };

    type Pos struct { x, y int };

    type Direction_And_Pos struct {
        pos Pos;
        direction Direction;
    }
    directions := []Direction_And_Pos{
        {Pos{ 1,  0}, RIGHT},
        {Pos{-1,  0}, LEFT },
        {Pos{ 0, -1}, UP   },
        {Pos{ 0,  1}, DOWN },
    };

    type Pos_And_Prev struct { curr, last Pos };
    // using recursion, check every path we can possibly take,
    //
    // just don't go backwards.
    stack := make([]Pos_And_Prev, 0);
    // put the start on the stack
    // make the last pos an invalid position.
    //
    // this is so I dont have to duplicate code to start the loop.
    Append(&stack, Pos_And_Prev{Pos{0, 0}, Pos{-1, -1}});

    for len(stack) > 0 {
        item := Pop(&stack);
        curr, last := item.curr, item.last;

        path := grid[curr.y][curr.x]-1;
        path_connection := path_connections[path]

        // we reached the end!
        if curr.x == n-1 && curr.y == m-1 { return true; }

        if been_to[curr.y][curr.x] { continue; }
        been_to[curr.y][curr.x] = true;

        for _, dir := range directions {
            next := Pos{curr.x + dir.pos.x, curr.y + dir.pos.y};
            // don't go backwards
            if next == last { continue; }
            // bounds check on grid
            if next.x < 0 || next.x > n-1 { continue; }
            if next.y < 0 || next.y > m-1 { continue; }

            // check if the road goes this way.
            if path_connection & dir.direction == 0 { continue; }
            { // check if the other way is also fine.
                next_path := grid[next.y][next.x]-1;
                next_connections := path_connections[next_path];
                // kinda funny trick, because of the way we set up our const's
                // the opposite direction is 2 shifts in one direction, check both.
                connected := false;
                connected = connected || (next_connections & (dir.direction << 2) != 0);
                connected = connected || (next_connections & (dir.direction >> 2) != 0);
                if !connected { continue; }
            }


            // good next position
            Append(&stack, Pos_And_Prev{next, curr});
        }

    }

    return false;
}


func Append[T any](slice *[]T, items ...T) {
	*slice = append(*slice, items...);
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
