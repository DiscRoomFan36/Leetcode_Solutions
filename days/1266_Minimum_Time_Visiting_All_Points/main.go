package main

func minTimeToVisitAllPoints(points [][]int) int {
    distance := 0;
    for i := range len(points)-1 {
        x1, y1 := points[i  ][0], points[i  ][1];
        x2, y2 := points[i+1][0], points[i+1][1];

        // Chebyshev distance
        distance += max(Abs(x1-x2), Abs(y1-y2));
    }
    return distance;
}

func Abs(x int) int {
    if x < 0 { x = -x; }
    return x;
}