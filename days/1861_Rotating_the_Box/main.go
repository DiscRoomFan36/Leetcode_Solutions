package main

func rotateTheBox(boxGrid [][]byte) [][]byte {
    if len(boxGrid) == 0 { return [][]byte{}; }
    m, n := len(boxGrid), len(boxGrid[0]);

    new_box := make_grid[byte](m, n);


    rot_90 := func(x, y int) (int, int) {
        // this is wrong.
        return m-y-1, x;
    }

    // copy into rotated box.
    for y := range m {
        for x := range n {
            new_x, new_y := rot_90(x, y);
            new_box[new_y][new_x] = boxGrid[y][x];
        }
    }

    // do sinulation.
    for x := range m {
        place_to_land := n-1;
        for y := n-1; y >= 0; y-- {
            cell := new_box[y][x];
            switch cell {
                case '.': { /* do nothing */ }
                case '*': {
                    // new floor
                    place_to_land = y-1;
                }
                case '#': {
                    // maybe fall
                    if place_to_land == y {
                        // this is the new place to land.
                        place_to_land = y-1;
                    } else {
                        // fall
                        new_box[place_to_land][x] = '#';
                        new_box[y][x] = '.';
                        place_to_land -= 1;
                    }
                }
                default: panic("not possible.");
            }
        }
    }

    return new_box;
}

// makes a grid, with default values. i hope go can turn this into something cool.
//
// might be better to make([]T, m*n) then grab some slices.
func make_grid[T any](n, m int) [][]T {
    grid := make([][]T, m);
    for j := range m { grid[j] = make([]T, n); }
    return grid;
}
