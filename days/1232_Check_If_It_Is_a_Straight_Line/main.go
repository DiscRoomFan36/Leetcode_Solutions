package main

func checkStraightLine(coordinates [][]int) bool {
    n := len(coordinates)
    if n <= 2 { return true }

    // rise over run
    x1, y1 := coordinates[0][0], coordinates[0][1]
    x2, y2 := coordinates[1][0], coordinates[1][1]
    for i := 2; i < n; i++ {
        x3, y3 := coordinates[i][0], coordinates[i][1]

        x12 := x2 - x1
        y12 := y2 - y1

        x13 := x3 - x1
        y13 := y3 - y1

        if  y12 * x13 != y13 * x12 { return false }
    }

    return true
}
