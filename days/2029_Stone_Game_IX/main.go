package main

func stoneGameIX(stones []int) bool {
    mod_counts := [3]int{};
    for _, stone := range stones {
        mod_counts[stone % 3] += 1;
    }

    if mod_counts[0] % 2 == 0 {
        // alice can chose the one that wins, only if there are enought not to run out of stones
        return mod_counts[1] >= 1 && mod_counts[2] >= 1;
    }

    // else, its kinda flipped, alice can still choose,
    // but there must be enought of a difference in piles.
    return Abs(mod_counts[1] - mod_counts[2]) > 2;
}

func Abs(x int) int {
    if x < 0 { x = -x; }
    return x;
}
