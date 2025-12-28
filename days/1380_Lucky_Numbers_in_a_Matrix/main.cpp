#include <stdlib.h>
#include <stdio.h>

#include <stddef.h>
#include <stdint.h>
#include <stdbool.h>

#include <unordered_map>
#include <vector>

using std::vector;
using std::unordered_map;


class Solution {
public:
    vector<int> luckyNumbers(vector<vector<int>>& matrix) {
        size_t m = matrix.size();
        size_t n = matrix[0].size();

        int min_in_rows[m];
        int max_in_cols[n];

        for (size_t i = 0; i < m; i++) min_in_rows[i] = 100000;
        // where are you memset?
        for (size_t i = 0; i < n; i++) max_in_cols[i] = 0;

        for (size_t j = 0; j < m; j++) {
            for (size_t i = 0; i < n; i++) {
                int x = matrix[j][i];

                if (x < min_in_rows[j]) min_in_rows[j] = x;
                if (x > max_in_cols[i]) max_in_cols[i] = x;
            }
        }


        for (size_t j = 0; j < m; j++) {
            for (size_t i = 0; i < n; i++) {
                int x = matrix[j][i];

                if (x == min_in_rows[j] && x == max_in_cols[i]) {
                    // there can only be one luck number,
                    return {x};
                }
            }
        }
        return {};
    }
};

int main(void) {
    printf("hello world!\n");

    return 0;
}
