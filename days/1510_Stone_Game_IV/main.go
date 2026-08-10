package main

func winnerSquareGame(n int) bool {
    i_win := make([]bool, n+1);
    for i := 1; i <= n; i++ {
        can_make_oppent_lose := false;
        for j := 1; j*j <= i; j++ {
            if i_win[i - j*j] == false {
                can_make_oppent_lose = true;
                break;
            }
        }
        i_win[i] = can_make_oppent_lose;
    }
    return i_win[n];
}
