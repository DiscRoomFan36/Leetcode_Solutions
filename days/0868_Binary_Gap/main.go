package main

func binaryGap(n int) int {
    max_distance := 0;

    last_bit_index := -1;
    for i := range 64 {
        // dont check every bit.
        if n < (1 << i) { break; }

        if ((n >> i) & 1) != 0 {
            if last_bit_index != -1 {
                dist := i - last_bit_index;
                max_distance = max(max_distance, dist);
            }

            last_bit_index = i;
        }
    }

    return max_distance;
}