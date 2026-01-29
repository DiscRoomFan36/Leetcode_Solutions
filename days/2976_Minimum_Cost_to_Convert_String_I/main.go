package main

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
    const NUM_LETTERS = 26
    const NL = NUM_LETTERS

    const INF = 1 << 60

    // oh merciful gods above! grant me a 1d array!
    edges := [NL][NL]int{}
    for j := range NL { for i := range NL { edges[j][i] = INF } }

    for i := range original {
        from, to, c := original[i] - 'a', changed[i] - 'a', cost[i]
        // the authors are being dickheads.
        edges[from][to] = min(edges[from][to], c)
    }

    // Floyd-Warshall algorithm
    // 26*26*26 = 17576
    for k := range NL {
        for i := range NL {
            for j := range NL {
                edges[i][j] = min(edges[i][j], edges[i][k] + edges[k][j])
            }
        }
    }

    result := 0
    for i := range len(source) {
        a, b := source[i] - 'a', target[i] - 'a'
        // we don't need to do anything
        if a == b { continue }

        weight := edges[a][b]
        if weight == INF { return -1 }
        result += weight
    }
    return int64(result)
}