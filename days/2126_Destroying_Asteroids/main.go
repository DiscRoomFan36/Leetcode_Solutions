package main

import "slices"

func asteroidsDestroyed(mass int, asteroids []int) bool {
    slices.Sort(asteroids);

    for _, asteroid := range asteroids {
        if mass < asteroid { return false; }
        mass += asteroid;
    }
    return true;
}
