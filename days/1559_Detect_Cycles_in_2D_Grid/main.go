package main

func containsCycle(grid [][]byte) bool {
    if len(grid) == 0 { return false; }
    m, n := len(grid), len(grid[0]);

    // the positions weve been too, to help find when we are in a cycle.
    been_to := make_grid[bool](n, m);

    // assumes you havent been_to this position before.
    //
    // checks for a cycle, and marks cells as been_to,
    path_contains_cycle := func(x, y int) bool {

        // get the correct char we work with.
        //
        // good to grab it just once here.
        correct_char := grid[y][x];

        type Pos          struct { x, y int };
        type Pos_And_Prev struct { curr, last Pos };

        // using recursion, check every path we can possibly take,
        //
        // just don't go backwards.
        stack := make([]Pos_And_Prev, 0);
        // put the current position on the stack,
        // make the last one an invalid position.
        //
        // this is so I dont have to duplicate code to start the loop.
        Append(&stack, Pos_And_Prev{Pos{x, y}, Pos{-1, -1}});

        directions := []Pos{
            { 1,  0},
            {-1,  0},
            { 0,  1},
            { 0, -1},
        };

        // stack based recursion.
        for len(stack) > 0 {
            item := Pop(&stack);
            curr, last := item.curr, item.last;

            // if we have been here before, its a cycle.
            if been_to[curr.y][curr.x] { return true; }
            been_to[curr.y][curr.x] = true;

            for _, dir := range directions {
                next := Pos{curr.x + dir.x, curr.y + dir.y};
                // cant go backwards
                if next == last { continue; }
                // bounds check on grid
                if next.x < 0 || next.x > n-1 { continue; }
                if next.y < 0 || next.y > m-1 { continue; }
                // finally check if its the right char
                if grid[next.y][next.x] != correct_char { continue; }

                // good next position
                Append(&stack, Pos_And_Prev{next, curr});
            }
        }

        // we detected no cycle
        return false;
    }

    // now just find every cell that we havent been_to, and check if its a cycle.
    for j := range m {
        for i := range n {
            if !been_to[j][i] {
                if path_contains_cycle(i, j) {
                    return true;
                }
            }
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

