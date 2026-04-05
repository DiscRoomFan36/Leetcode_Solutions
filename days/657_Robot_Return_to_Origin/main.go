package main

func judgeCircle(moves string) bool {
    if len(moves) % 2 != 0 { return false; }
    U_count := 0;
    D_count := 0;
    L_count := 0;
    R_count := 0;
    for _, c := range moves {
        if c == 'U' { U_count += 1; }
        if c == 'D' { D_count += 1; }
        if c == 'L' { L_count += 1; }
        if c == 'R' { R_count += 1; }
    }
    return U_count == D_count && L_count == R_count;
}
