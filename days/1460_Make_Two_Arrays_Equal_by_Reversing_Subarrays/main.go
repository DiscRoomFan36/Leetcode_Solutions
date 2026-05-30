package main

// this is just "do these array countain the same elements"
func canBeEqual(target []int, arr []int) bool {
    if len(target) != len(arr) { return false; }
    
    counts := make(map[int]int);
    for i := range len(arr) {
        counts[arr   [i]] += 1;
        counts[target[i]] -= 1;
    }

    for _, count := range counts {
        if count != 0 { return false; }
    }
    return true;
}
