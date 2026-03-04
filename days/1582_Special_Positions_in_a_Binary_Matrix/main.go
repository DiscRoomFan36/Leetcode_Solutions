package main

func numSpecial(mat [][]int) int {
    if len(mat) == 0 { return 0; }
    m, n := len(mat), len(mat[0]);

    // preprocess to make checking for only one in col faster.
    col_sums := make([]int, n);
    for j := range m {
        for i := range n {
            col_sums[i] += mat[j][i];
        }
    }


    result := 0;
    for j := range m {
        // check every row to see if there is only one
        // in the row, and where it is.
        row_sum := 0;
        index_of_one := 0;
        for i := range n {
            item := mat[j][i];

            if item == 1 { index_of_one = i; }
            row_sum += item;
        }

        if row_sum == 1 {
            if col_sums[index_of_one] == 1 {
                result += 1;
            }
        }
    }
    return result;
}