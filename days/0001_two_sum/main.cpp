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
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int, int> map = {};
        for (size_t i = 0; i < nums.size(); i++) {
            int compliment = target - nums[i];
            if (map.count(compliment)) {
                return {map[compliment], (int)i};
            }
            map[nums[i]] = i;
        }
        return {};
    }
};


int main(void) {
    printf("hello world!\n");

    return 0;
}
