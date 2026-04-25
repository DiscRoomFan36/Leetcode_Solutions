package main

import (
	"cmp"
	"slices"
)

func maxDistance(side int, _points [][]int, k int) int {
    if len(_points) == 0 { return 0; }

    // sort the points so they are in a clockwise order.
    //
    // make function "can get length n"
    //
    // binary search

    // sort points in clockwise order
    type Point struct { x, y int };

    points := make([]Point, len(_points));
    for i := range _points {
        points[i].x = _points[i][0];
        points[i].y = _points[i][1];
    }

    edge_number := func(p Point) int {
        if        p.x == 0    { return 0;
        } else if p.y == side { return 1;
        } else if p.x == side { return 2;
        } else                { return 3; }
    }

    slices.SortFunc(points,
        func(a, b Point) int {
            a_edge := edge_number(a);
            b_edge := edge_number(b);

            if a_edge != b_edge {
                return cmp.Compare(a_edge, b_edge);
            }

            // the edge is the same
            if a_edge <= 1 {
                // on one edge, one component (either x or y) is the
                // same for both of them. sort by the other one.
                return cmp.Compare(a.x + a.y, b.x + b.y);
            } else {
                return cmp.Compare(b.x + b.y, a.x + a.y);
            }
        },
    );

    manhatten_distance := func(a, b Point) int {
        return Abs(a.x - b.x) + Abs(a.y - b.y);
    }

    // make function "can get length n"
    can_get_min_distance_of_n := func(n int) bool {
        //
        // hack, but this will handle the case where the 0th point is not the best one.
        //
        // 800ms time, but still a win.
        //
        // the faster way would be to find the exact range where we need to check
        // the only solutions would be in the range [0..index_of_first_node_that_is_far_enough_from_the_0th_node)
        //
        for start_index := 0; start_index < min(1000, len(points)); start_index += 1 {
            start_point := points[start_index % len(points)];

            points_left_to_add := k-1; // because the first point is allready there
            last_point := start_point;

            for i := 0; i < len(points); i += 1 {
                new_point := points[(start_index + i) % len(points)];
                if manhatten_distance(last_point, new_point) >= n {
                    last_point = new_point;
                    points_left_to_add -= 1;
                    // got enough points.
                    if points_left_to_add == 0 {
                        // check the last point is far enough from the first point.
                        if manhatten_distance(last_point, start_point) >= n {
                            return true;
                        } else {
                            break; // maybe check another start index.
                        }
                    }
                }
            }
        }

        return false;
    };

    // binary search on maximum possible distance.
    //
    // because min(k) == 4, this is true.but we could
    // make this side * 2, to be able to handle min(k) == 2
    max_possible_distance := side;

    best_size := 1;

    low, high := 1, max_possible_distance;
    for low <= high {
        mid := (low + high) / 2;
        if can_get_min_distance_of_n(mid) {
            best_size = mid;
            low  = mid+1;
        } else {
            high = mid-1; 
        }
    }

    return best_size;
}

func Abs(x int) int {
    if x < 0 { x *= -1; }
    return x;
}
