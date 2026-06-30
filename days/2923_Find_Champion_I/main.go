package main

func findChampion(grid [][]int) int {
    n := len(grid);

    someone_is_stronger := make([]bool, n);
    for j := range n {
        // we only need to search one half of the grid.
        for i := range j {
            j_beats_i := grid[j][i] == 1;

            if j_beats_i {
                someone_is_stronger[i] = true;
            } else {
                someone_is_stronger[j] = true;
            }
        }
    }

    for i, anyone_stronger := range someone_is_stronger {
        if !anyone_stronger { return i; }
    }
    panic("no clear winner");
}
