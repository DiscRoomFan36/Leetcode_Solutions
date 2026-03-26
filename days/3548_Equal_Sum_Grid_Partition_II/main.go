package main

type Pos struct { x, y int; };

func canPartitionGrid(grid [][]int) bool {
    if len(grid) == 0 { return false; }
    m, n := len(grid), len(grid[0]);

    prefix_sum_grid := make_grid[int](n+1, m+1);

    num_to_position := make(map[int][]Pos);

    for j := range m {
        for i := range n {
            num := grid[j][i];

            up      := prefix_sum_grid[j  ][i+1];
            left    := prefix_sum_grid[j+1][i  ];
            up_left := prefix_sum_grid[j  ][i  ];
            prefix_sum_grid[j+1][i+1] = (up + left - up_left) + num;

            num_to_position[num] = append(num_to_position[num], Pos{ i, j });
        }
    }

    // includes the two end points.
    get_area_sum := func(x1, y1, x2, y2 int) int {
        w := x2 - x1 + 1;
        h := y2 - y1 + 1;
        return prefix_sum_grid[y1+h][x1+w] - prefix_sum_grid[y1+h][x1] - prefix_sum_grid[y1][x1+w] + prefix_sum_grid[y1][x1];
    }

    for i := range n-1 {
        sum_1 := get_area_sum(0, 0,     i, m-1);
        sum_2 := get_area_sum(i+1, 0, n-1, m-1);
        if sum_1 == sum_2 { return true; }

        // else try and remove something to balance.
        w1, w2 := i+1, n-i-1;
        h := m;
        if sum_1 < sum_2 {
            num := sum_2 - sum_1;
            positions := num_to_position[num];

            for _, p := range positions {
                // is on the right side.
                on_right_side_of_line := p.x > i;
                blocking_path := (w2 == 1 && is_between(p.y, 0, m-1)) || (h == 1 && is_between(p.x, i+1, n-1));
                if on_right_side_of_line && !blocking_path {
                    return true;
                }
            }
        } else {
            num := sum_1 - sum_2;
            positions := num_to_position[num];

            for _, p := range positions {
                // is on the right side.
                on_right_side_of_line := p.x <= i;
                blocking_path := (w1 == 1 && is_between(p.y, 0, m-1)) || (h == 1 && is_between(p.x, 0, i));
                if on_right_side_of_line && !blocking_path {
                    return true;
                }
            }
        }
    }
    for j := range m-1 {
        sum_1 := get_area_sum(0, 0,   n-1,   j);
        sum_2 := get_area_sum(0, j+1, n-1, m-1);
        if sum_1 == sum_2 { return true; }

        // else try and remove something to balance.
        h1, h2 := j+1, m-j-1;
        w := n;
        if sum_1 < sum_2 {
            num := sum_2 - sum_1;
            positions := num_to_position[num];

            for _, p := range positions {
                // is on the right side.
                on_right_side_of_line := p.y > j;
                blocking_path := (h2 == 1 && is_between(p.x, 0, n-1)) || (w == 1 && is_between(p.y, j+1, m-1));
                if on_right_side_of_line && !blocking_path {
                    return true;
                }
            }
        } else {
            num := sum_1 - sum_2;
            positions := num_to_position[num];

            for _, p := range positions {
                // is on the right side.
                on_right_side_of_line := p.y <= j;
                blocking_path := (h1 == 1 && is_between(p.x, 0, n-1)) || (w == 1 && is_between(p.y, 0, j));
                if on_right_side_of_line && !blocking_path {
                    return true;
                }
            }
        }
    }

    return false;
}

// makes a grid, with default values. i hope go can turn this into something cool.
//
// might be better to make([]T, m*n) then grab some slices.
func make_grid[T any](n, m int) [][]T {
    grid := make([][]T, m)
    for j := range m { grid[j] = make([]T, n) }
    return grid
}

func abs(x int) int {
    if x < 0 { x *= -1; }
    return x;
}

func is_between(x, low, high int) bool {
    return low < x && x < high;
}
