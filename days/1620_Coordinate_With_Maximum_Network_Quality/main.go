package main

import "math"

func bestCoordinate(towers [][]int, radius int) []int {
    type Vec2 struct { x, y int; }
    best_pos := Vec2{0, 0};
    best_strength := -1;

    for x := range 50+1 {
        for y := range 50+1 {
            this_pos := Vec2{x, y};
            total_strength := 0;

            for _, tower := range towers {
                tower_x, tower_y, strength := tower[0], tower[1], tower[2];

                dist_sqr := Euclidean_distance(this_pos.x, this_pos.y, tower_x, tower_y);
                if dist_sqr > radius*radius { continue; }

                quality := float64(strength) / (1 + math.Sqrt(float64(dist_sqr)));
                total_strength += int(math.Floor(quality));
            }

            if total_strength > best_strength {
                best_strength = total_strength;
                best_pos = Vec2{x, y};
            }
        }
    }

    return []int{best_pos.x, best_pos.y};
}

func Euclidean_distance(x1, y1, x2, y2 int) int {
    x := x1 - x2;
    y := y1 - y2;
    return x*x + y*y;
}
