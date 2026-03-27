package main

func areSimilar(mat [][]int, k int) bool {
    if len(mat) == 0 { return true; }
    m, n := len(mat), len(mat[0]);

    // if we shift more than n times, we get the same result.
    k %= n;
    if k == 0 { return true; }

    for j := range m {
        for i := range n {
            index := (i + k) % n;
            if j % 2 == 0 {
                // even rows shift left.
                // proper mod.
                index = (((i - k) % n) + n) % n;
            }
            if mat[j][i] != mat[j][index] {
                return false;
            }
        }
    }

    return true;
}
