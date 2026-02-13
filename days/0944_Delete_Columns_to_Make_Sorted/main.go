package main

func minDeletionSize(strs []string) int {
    if len(strs) == 0 { return 0 }

    m, n := len(strs), len(strs[0])
    // fmt.Println(m, n)

    result := 0
    for i := range n {
        sorted := true
        for j := 1; j < m; j++ {
            if strs[j-1][i] > strs[j][i] {
                sorted = false
                break
            }
        }
        if !sorted { result += 1 }
    }
    return result
}