package main

func findChampion(n int, edges [][]int) int {
    someone_is_stronger := make([]bool, n);
    
    for _, edge := range edges {
        _, v := edge[0], edge[1];
        someone_is_stronger[v] = true;
    }

    winner := -1;
    for i, anyone_stronger := range someone_is_stronger {
        if !anyone_stronger {
            if winner != -1 { return -1; }
            winner = i;
        }
    }
    if winner == -1 {
        panic("no clear winner");
    }
    return winner;
}
