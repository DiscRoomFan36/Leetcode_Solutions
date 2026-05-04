package main

func rotate(matrix [][]int)  {
    n := len(matrix);

    next_position := func(x, y int) (int, int) {
        return n-y-1, x;
    }

    for y := range n {
        for x := y; x < n-y-1; x++ {

            // rotate all 4 positions.
            right_x, right_y := next_position(      x,       y);
            down_x,  down_y  := next_position(right_x, right_y);
            left_x,  left_y  := next_position( down_x,  down_y);

            up    := &matrix[      y][      x];
            right := &matrix[right_y][right_x];
            down  := &matrix[ down_y][ down_x];
            left  := &matrix[ left_y][ left_x];

            // swap 4 in a circle.
            *up, *right, *down, *left = *left, *up, *right, *down;
        }
    }
}
