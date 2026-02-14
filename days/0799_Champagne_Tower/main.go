package main

func champagneTower(poured int, query_row int, query_glass int) float64 {
    const NUM_ROWS_OF_WINE = 100
    // need an extra thing of wine, so it dosnt overflow to the right
    current := [NUM_ROWS_OF_WINE+1]float64{}
    next    := [NUM_ROWS_OF_WINE+1]float64{}

    // pour the inital wine.
    current[0] = float64(poured)

    // loop until we get to the required row.
    for j := range query_row {
        // only need to check the glasses that might have wine in them
        for i := range j+1 {
            // this branch hopefully goes away
            if current[i] > 1 {
                wine_spilled := current[i]-1 // one stays in the glass
                next[i  ] += wine_spilled/2
                next[i+1] += wine_spilled/2
                // current[i] = 1 // current will be stomped at the end of the turn.
            }
        }

        // ready for next loop, might be faster to 
        current = next
        // clear the next line
        next = [NUM_ROWS_OF_WINE+1]float64{}
    }

    // if its over 1, the rest has spilled to the bottom.
    return min(current[query_glass], 1)
}