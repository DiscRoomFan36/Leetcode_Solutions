package main

func maximalRectangle(matrix [][]byte) int {
    if len(matrix) == 0 { return 0 }
    m, n := len(matrix), len(matrix[0])

    // how far right can we go while only seeing 1's
    right_space := make([]int, m*n)
    // how far  down can we go while only seeing 1's
    down_space  := make([]int, m*n)

    // precalculate how far right and down you can go from any tile.
    for j := m-1; j >= 0; j-- {
        for i := n-1; i >= 0; i-- {
            is_one := (matrix[j][i] == '1')
            index_2d := j * n + i

            if !is_one {
                right_space[index_2d] = 0
                down_space [index_2d] = 0
                continue
            }

            {
                val_right := 0
                if i != n-1 { val_right = right_space[j*n + i+1] }
                right_space[index_2d] = val_right + 1
            }

            {
                val_down := 0
                if j != m-1 { val_down = down_space[(j+1)*n + i] }
                down_space[index_2d] = val_down + 1
            }
        }
    }

    best_size := 0
    for j := range m {
        for i := range n {
            index_2d := j * n + i
            right := right_space[index_2d]
            down  := down_space [index_2d]

            // this might not be true, but it is a fast check, plus, if its zero, it will exit here.
            best_size_this_could_possibly_make := right * down
            if best_size_this_could_possibly_make <= best_size { continue }

            min_down := 999999999999999
            // now find the for real biggest rect.
            for k := 0; k < right; k++ {
                min_down = min(min_down, down_space[index_2d+k])

                rec_size := (k+1) * min_down
                best_size = max(best_size, rec_size)

                // if the best possible isnt better, quit early
                if min_down * right <= best_size { break }
            }
        }
    }

    return best_size
}