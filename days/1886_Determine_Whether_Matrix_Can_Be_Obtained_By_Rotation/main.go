package main

func findRotation(mat [][]int, target [][]int) bool {
    if len(mat) == 0 { panic("leetcode problem"); }
    n := len(mat);

    type Vec2i struct { x, y int; }

    correct_rotation := [4]bool{true, true, true, true};
    for j := range n {
        for i := range n {
            mat_num := mat[j][i];

            rotations := [4]Vec2i{
                { x:     i,     y:     j     },
                { x: n - j - 1, y:     i     },
                { x: n - i - 1, y: n - j - 1 },
                { x:     j,     y: n - i - 1 },
            };

            for k, pos := range rotations {
                is_correct := (target[pos.y][pos.x] == mat_num);
                correct_rotation[k] = correct_rotation[k] && is_correct;
            }
        }
    }

    one_is_true := false;
    for _, correct := range correct_rotation {
        one_is_true = one_is_true || correct;
    }
    return one_is_true;
}
