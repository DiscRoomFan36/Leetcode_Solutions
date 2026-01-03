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
private:
    typedef uint64_t u64;

    static u64 pow_mod(u64 x, u64 y, u64 mod) {
        if (x == 0) return 0;

        u64 result = 1;
        u64 keep_squaring = x;
        while (y != 0) {
            if (y & 1) result = (result * keep_squaring) % mod;
            y = y >> 1;
            keep_squaring = (keep_squaring*keep_squaring) % mod;
        }

        return result;
    }


public:
    int countGoodNumbers(long long n) {
        const u64 MOD = 1000000007;

        const u64 NUM_GOOD_EVEN = 5; // 0, 2, 4, 6, 8
        const u64 NUM_GOOD_ODD  = 4; // 2, 3, 5, 7

        u64 even = pow_mod(NUM_GOOD_EVEN, (n + 1) / 2, MOD);
        u64 odd  = pow_mod(NUM_GOOD_ODD,  (n    ) / 2, MOD);

        return (even * odd) % MOD;
    }
};



int main(void) {
    printf("hello world!\n");

    return 0;
}
