package main

func rangeAddQueries(n int, queries [][]int) [][]int {
    matrix := make_grid[int](n, n);

    type Query struct {
        x1, y1, x2, y2 int;
    }

    // group like queries.
    query_map := make(map[Query]int)
    for k := range queries {
        y1, x1, y2, x2 := queries[k][0], queries[k][1], queries[k][2], queries[k][3];

        new_query := Query{x1, y1, x2, y2};
        query_map[new_query] += 1;
    }

    for query, count := range query_map {
        x1, y1, x2, y2 := query.x1, query.y1, query.x2, query.y2;

        for j := y1; j <= y2; j++ {
            for i := x1; i <= x2; i++ {
                matrix[j][i] += count;
            }
        }
    }
    return matrix;
}

// makes a grid, with default values. i hope go can turn this into something cool.
//
// might be better to make([]T, m*n) then grab some slices.
func make_grid[T any](n, m int) [][]T {
    grid := make([][]T, m)
    for j := range m { grid[j] = make([]T, n) }
    return grid
}
