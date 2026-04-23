package main

func distance(nums []int) []int64 {

    // the real problem they are going to hit me with is,
    //
    // array of ints
    // [1, 4, 7, 9, 13]
    //
    // find abs between all.

    number_to_indexes := make(map[int][]int);
    for i, n := range nums {
        number_to_indexes[n] = append(number_to_indexes[n], i);
    }

    result := make([]int64, len(nums));

    for _, indexes := range number_to_indexes {
        // [1, 4, 7, 9, 13]
        //
        // 7
        // (7 - 4) + (7 - 1) + (9 - 7) + (13 - 7)

        // smart O(n)

        prefix_forward  := make([]int, len(indexes)+1);
        prefix_backward := make([]int, len(indexes)+1);

        for i := range indexes {
            j := len(indexes) - i - 1;

            prefix_forward [i+1] = prefix_forward [i  ] + indexes[i];
            prefix_backward[j  ] = prefix_backward[j+1] + indexes[j];
        }

        for i, index := range indexes {
            distances_before := (index * i) - prefix_forward[i]
            distances_after  := prefix_backward[i+1] - (index * (len(indexes)-i-1))
            result[index] = int64(distances_before + distances_after);
        }
    }

    return result;
}

func Abs(x int) int {
    if x < 0 { x *= -1; }
    return x;
}
