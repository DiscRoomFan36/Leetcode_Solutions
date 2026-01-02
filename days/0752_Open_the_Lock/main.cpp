#include <stdlib.h>
#include <stdio.h>

#include <stddef.h>
#include <stdint.h>
#include <stdbool.h>

#include <unordered_map>
#include <vector>
#include <string>

using std::vector;
using std::unordered_map;
using std::string;

class Solution {
public:
    int openLock(vector<string>& deadends, string target) {
        int distances[10000] = {0};

        int start = 0;
        int end = stoi(target);

        if (start == end) return 0;

        distances[start] = 1;

        for (auto &&deadend : deadends) {
            int num = stoi(deadend);
            if (num == start) return -1;
            distances[num] = -1;
        }


        vector<int> positions_1 = {start};
        vector<int> positions_2 = {};

        auto *positions       = &positions_1;
        auto *other_positions = &positions_2;

        while (positions->size()) {
            other_positions->clear();

            // i hate c++
            for (auto &&pos : *positions) {
                int dist = distances[pos];

                // we need all other positions we can reach

                int d1 = (pos       ) % 10;
                int d2 = (pos /   10) % 10;
                int d3 = (pos /  100) % 10;
                int d4 = (pos / 1000) % 10;

                int d1p = pos + ((((d1 + 1)     ) % 10) - d1) *    1;
                int d1m = pos + ((((d1 - 1) + 10) % 10) - d1) *    1;

                int d2p = pos + ((((d2 + 1)     ) % 10) - d2) *   10;
                int d2m = pos + ((((d2 - 1) + 10) % 10) - d2) *   10;

                int d3p = pos + ((((d3 + 1)     ) % 10) - d3) *  100;
                int d3m = pos + ((((d3 - 1) + 10) % 10) - d3) *  100;

                int d4p = pos + ((((d4 + 1)     ) % 10) - d4) * 1000;
                int d4m = pos + ((((d4 - 1) + 10) % 10) - d4) * 1000;

                int options[8] = {d1p, d1m, d2p, d2m, d3p, d3m, d4p, d4m};

                for (size_t i = 0; i < 8; i++) {
                    // printf("options[%zu] = %d\n", i, options[i]);

                    int option = options[i];
                    // check if were at the end
                    if (option == end) return dist;

                    if (distances[option] == -1) {
                        continue; // was a dead end
                    } else if (distances[option] == 0) {
                        // new position
                        distances[option] = dist + 1;
                        other_positions->push_back(option);
                    } else {
                        // we were allready there
                    }
                }
            }

            auto tmp = positions;
            positions = other_positions;
            other_positions = tmp;
        }

        return -1;
    }
};



int main(void) {
    printf("hello world!\n");

    return 0;
}
