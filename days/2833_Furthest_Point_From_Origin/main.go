package main

func furthestDistanceFromOrigin(moves string) int {
    l_count, r_count, under_count := 0, 0, 0;
    for _, c := range moves {
        switch c {
            case 'L': l_count     += 1;
            case 'R': r_count     += 1;
            case '_': under_count += 1;
        }
    }

    return Abs(l_count - r_count) + under_count;
}

func Abs(x int) int {
    if x < 0 { x *= -1; }
    return x;
}
