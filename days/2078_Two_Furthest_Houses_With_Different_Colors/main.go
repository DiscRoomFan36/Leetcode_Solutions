package main

func maxDistance(colors []int) int {
    best_distance := 1; // there will always be one.
    for i, color_1 := range colors {
        // will always be smaller.
        if i != 0 && color_1 == colors[i-1] { continue; }

        for j := i+best_distance; j < len(colors); j++ {
            color_2 := colors[j];

            if color_1 != color_2 {
                best_distance = max(best_distance, j-i);
            }
        }
    }
    return best_distance;
}
