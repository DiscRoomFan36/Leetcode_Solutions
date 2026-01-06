
package main


func numberOfArrays(differences []int, lower int, upper int) int {
    total_range := upper - lower + 2

    min_diff := 0
    max_diff := 0
    running  := 0
    for _, n := range differences {
        running += n
        if min_diff > running { min_diff = running }
        if max_diff < running { max_diff = running }

        if max_diff - min_diff + 1 >= total_range { return 0 }
    }

    diff_range  := max_diff - min_diff + 1
    return total_range - diff_range
}
