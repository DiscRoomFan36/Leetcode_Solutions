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
    int findSubstringInWraproundString(string s) {
        int counts[26] = {0};
        // small special case for first char.
        // 1 <= s.length <= 10^5, this is safe
        counts[s[0] - 'a'] += 1;

        int max_length_current = 1;
        for (size_t i = 1; i < s.length(); i++) {
            char c1 = s[i-1];
            char c2 = s[i];
            if ((c1+1 == c2) || (c1 == 'z' && c2 == 'a')) {
                max_length_current += 1;
            } else {
                max_length_current = 1;
            }

            int index = c2 - 'a';
            if (counts[index] < max_length_current) {
                counts[index] = max_length_current;
            }
        }

        int total = 0;
        for (size_t i = 0; i < sizeof(counts) / sizeof(counts[0]); i++) {
            total += counts[i];
        }
        return total;
    }
};

int main(void) {
    printf("hello world!\n");

    return 0;
}
